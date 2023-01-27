package gql

import (
	"context"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"github.com/pulumi/pulumi-go-provider"
	"github.com/zeet-dev/pulumi-zeet-native/provider/pkg/model"
	"net/http"
	"time"
)

type ZeetGraphqlClient struct {
	endpoint string
	apiToken string
	client   graphql.Client
}

func NewZeetGraphqlClient(endpoint string, apiToken string) ZeetGraphqlClient {
	return ZeetGraphqlClient{
		endpoint: endpoint,
		apiToken: apiToken,
		client: graphql.NewClient(endpoint,
			&http.Client{Transport: &authedTransport{apiToken: apiToken, wrapped: http.DefaultTransport}}),
	}
}

type authedTransport struct {
	apiToken string
	wrapped  http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.apiToken)
	return t.wrapped.RoundTrip(req)
}

func (c *ZeetGraphqlClient) GetCurrentUserID(ctx context.Context) (string, error) {
	resp, err := currentUserID(ctx, c.client)
	if err != nil {
		return "", err
	}
	return resp.CurrentUser.Id, nil
}

type CreateProjectResponse struct {
	ID        string
	Name      string
	UpdatedAt time.Time
}

func (c *ZeetGraphqlClient) CreateProject(ctx context.Context, userID string, name string) (CreateProjectResponse, error) {
	resp, err := createProject(ctx, c.client, userID, name)
	if err != nil {
		return CreateProjectResponse{}, err
	}
	project := resp.CreateProjectV2
	return CreateProjectResponse{
		ID:        project.Id,
		Name:      project.Name,
		UpdatedAt: project.UpdatedAt,
	}, nil
}

func (c *ZeetGraphqlClient) ReadProject(ctx context.Context, projectID string) (CreateProjectResponse, error) {
	resp, err := getProjectByID(ctx, c.client, projectID)
	if err != nil {
		return CreateProjectResponse{}, err
	}
	return CreateProjectResponse{
		ID:        resp.Project.Id,
		Name:      resp.Project.Name,
		UpdatedAt: resp.Project.UpdatedAt,
	}, nil
}

func (c *ZeetGraphqlClient) UpdateProject(ctx context.Context, projectID string, name *string) (CreateProjectResponse, error) {
	resp, err := updateProject(ctx, c.client, projectID, name)
	if err != nil {
		return CreateProjectResponse{}, err
	}
	return CreateProjectResponse{
		ID:        resp.UpdateProjectV2.Id,
		Name:      resp.UpdateProjectV2.Name,
		UpdatedAt: resp.UpdateProjectV2.UpdatedAt,
	}, nil
}

func (c *ZeetGraphqlClient) DeleteProject(ctx provider.Context, projectID string) error {
	resp, err := deleteProject(ctx, c.client, projectID)
	// graphql error
	if err != nil {
		return err
	}
	// server returned false
	if !resp.GetDeleteProjectV2() {
		return fmt.Errorf("unable to delete project '%s'", projectID)
	}
	// server returned true
	return nil
}

type CreateEnvironmentResponse struct {
	ID        string
	Name      string
	ProjectID string
	UpdatedAt time.Time
}

func (c *ZeetGraphqlClient) CreateEnvironment(ctx provider.Context, projectID string, name string) (CreateEnvironmentResponse, error) {
	resp, err := createEnvironment(ctx, c.client, projectID, name)
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}
	environment := resp.CreateProjectEnvironment
	project := environment.GetProject()
	if project == nil {
		return CreateEnvironmentResponse{}, fmt.Errorf("expected to find project for environment '%s'", environment.GetId())
	}
	return CreateEnvironmentResponse{
		ID:        environment.GetId(),
		Name:      environment.GetName(),
		ProjectID: project.GetId(),
		UpdatedAt: environment.GetUpdatedAt(),
	}, nil
}

func (c *ZeetGraphqlClient) ReadEnvironment(ctx context.Context, projectID string, environmentID string) (CreateEnvironmentResponse, error) {
	resp, err := getProjectEnvironments(ctx, c.client, projectID)
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}
	proj := resp.GetProject()
	if proj == nil {
		return CreateEnvironmentResponse{}, fmt.Errorf("could not fetch project by id '%s'", projectID)
	}
	// find environment by id in array
	for _, environment := range proj.GetEnvironments() {
		if environment == nil {
			continue
		}
		if environment.GetId() == environmentID {
			return CreateEnvironmentResponse{
				ID:        environment.GetId(),
				Name:      environment.GetName(),
				ProjectID: proj.GetId(),
				UpdatedAt: environment.GetUpdatedAt(),
			}, nil
		}
	}
	// environment was not found in project
	return CreateEnvironmentResponse{}, fmt.Errorf("no environment with id '%s' found in project '%s'", environmentID, projectID)
}

func (c *ZeetGraphqlClient) UpdateEnvironment(ctx context.Context, environmentID string, name *string) (CreateEnvironmentResponse, error) {
	resp, err := updateEnvironment(ctx, c.client, environmentID, name)
	if err != nil {
		return CreateEnvironmentResponse{}, err
	}
	updatedEnv := resp.GetUpdateProjectEnvironment()
	return CreateEnvironmentResponse{
		ID:        updatedEnv.GetId(),
		Name:      updatedEnv.GetName(),
		UpdatedAt: updatedEnv.GetUpdatedAt(),
	}, nil
}

func (c *ZeetGraphqlClient) DeleteEnvironment(ctx provider.Context, environmentID string) error {
	resp, err := deleteEnvironment(ctx, c.client, environmentID)
	// graphql error
	if err != nil {
		return err
	}
	// server returned false
	if !resp.GetDeleteProjectEnvironment() {
		return fmt.Errorf("unable to delete environment '%s'", environmentID)
	}
	// server returned true
	return nil
}

type CreateAppResponse struct {
	ID        string
	UpdatedAt time.Time
}

func (c *ZeetGraphqlClient) CreateApp(ctx provider.Context, args model.CreateAppInput) (CreateAppResponse, error) {
	buildType := BuildType(args.Build.Type)
	deployTarget := DeployTarget(args.Deploy.DeployTarget)
	input := CreateProjectGitInput{
		UserID:        &args.UserID,
		ProjectID:     &args.ProjectID,
		EnvironmentID: &args.EnvironmentID,
		Name:          &args.Name,
		Build: &ProjectBuildInput{
			BuildType:      &buildType,
			DockerfilePath: &args.Build.DockerfilePath,
		},
		DeployTarget: &ProjectDeployInput{
			DeployTarget: deployTarget,
			ClusterID:    &args.Deploy.ClusterID,
		},
		Envs: environmentVariablesToRequestInput(args.EnvironmentVariables),
	}
	if args.GithubInput != nil {
		input.Url = args.GithubInput.Url
		input.ProductionBranch = &args.GithubInput.ProductionBranch
	} else {
		return CreateAppResponse{}, fmt.Errorf("must specify one app spec")
	}
	resp, err := createAppGit(ctx, c.client, &input)
	if err != nil {
		return CreateAppResponse{}, err
	}
	return CreateAppResponse{
		ID:        resp.CreateProjectGit.Id,
		UpdatedAt: resp.CreateProjectGit.UpdatedAt,
	}, nil
}

func environmentVariablesToRequestInput(variables []model.CreateAppEnvironmentVariableInput) []*EnvVarInput {
	out := []*EnvVarInput{}
	for _, variable := range variables {
		input := &EnvVarInput{
			Name:   variable.Name,
			Value:  variable.Value,
			Sealed: variable.Sealed,
		}
		out = append(out, input)
	}
	return out
}

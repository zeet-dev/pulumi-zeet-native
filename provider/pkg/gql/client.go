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
	return CreateEnvironmentResponse{
		ID:        environment.Id,
		Name:      environment.Name,
		ProjectID: environment.Project.Id,
		UpdatedAt: environment.UpdatedAt,
	}, nil
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

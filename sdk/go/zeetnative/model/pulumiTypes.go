// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package model

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateAppBuildInput struct {
	DockerfilePath *string `pulumi:"dockerfilePath"`
	Type           string  `pulumi:"type"`
}

// CreateAppBuildInputInput is an input type that accepts CreateAppBuildInputArgs and CreateAppBuildInputOutput values.
// You can construct a concrete instance of `CreateAppBuildInputInput` via:
//
//	CreateAppBuildInputArgs{...}
type CreateAppBuildInputInput interface {
	pulumi.Input

	ToCreateAppBuildInputOutput() CreateAppBuildInputOutput
	ToCreateAppBuildInputOutputWithContext(context.Context) CreateAppBuildInputOutput
}

type CreateAppBuildInputArgs struct {
	DockerfilePath pulumi.StringPtrInput `pulumi:"dockerfilePath"`
	Type           pulumi.StringInput    `pulumi:"type"`
}

func (CreateAppBuildInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppBuildInput)(nil)).Elem()
}

func (i CreateAppBuildInputArgs) ToCreateAppBuildInputOutput() CreateAppBuildInputOutput {
	return i.ToCreateAppBuildInputOutputWithContext(context.Background())
}

func (i CreateAppBuildInputArgs) ToCreateAppBuildInputOutputWithContext(ctx context.Context) CreateAppBuildInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppBuildInputOutput)
}

type CreateAppBuildInputOutput struct{ *pulumi.OutputState }

func (CreateAppBuildInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppBuildInput)(nil)).Elem()
}

func (o CreateAppBuildInputOutput) ToCreateAppBuildInputOutput() CreateAppBuildInputOutput {
	return o
}

func (o CreateAppBuildInputOutput) ToCreateAppBuildInputOutputWithContext(ctx context.Context) CreateAppBuildInputOutput {
	return o
}

func (o CreateAppBuildInputOutput) DockerfilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateAppBuildInput) *string { return v.DockerfilePath }).(pulumi.StringPtrOutput)
}

func (o CreateAppBuildInputOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppBuildInput) string { return v.Type }).(pulumi.StringOutput)
}

type CreateAppDeployInput struct {
	ClusterId    *string `pulumi:"clusterId"`
	DeployTarget string  `pulumi:"deployTarget"`
}

// CreateAppDeployInputInput is an input type that accepts CreateAppDeployInputArgs and CreateAppDeployInputOutput values.
// You can construct a concrete instance of `CreateAppDeployInputInput` via:
//
//	CreateAppDeployInputArgs{...}
type CreateAppDeployInputInput interface {
	pulumi.Input

	ToCreateAppDeployInputOutput() CreateAppDeployInputOutput
	ToCreateAppDeployInputOutputWithContext(context.Context) CreateAppDeployInputOutput
}

type CreateAppDeployInputArgs struct {
	ClusterId    pulumi.StringPtrInput `pulumi:"clusterId"`
	DeployTarget pulumi.StringInput    `pulumi:"deployTarget"`
}

func (CreateAppDeployInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppDeployInput)(nil)).Elem()
}

func (i CreateAppDeployInputArgs) ToCreateAppDeployInputOutput() CreateAppDeployInputOutput {
	return i.ToCreateAppDeployInputOutputWithContext(context.Background())
}

func (i CreateAppDeployInputArgs) ToCreateAppDeployInputOutputWithContext(ctx context.Context) CreateAppDeployInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppDeployInputOutput)
}

type CreateAppDeployInputOutput struct{ *pulumi.OutputState }

func (CreateAppDeployInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppDeployInput)(nil)).Elem()
}

func (o CreateAppDeployInputOutput) ToCreateAppDeployInputOutput() CreateAppDeployInputOutput {
	return o
}

func (o CreateAppDeployInputOutput) ToCreateAppDeployInputOutputWithContext(ctx context.Context) CreateAppDeployInputOutput {
	return o
}

func (o CreateAppDeployInputOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateAppDeployInput) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o CreateAppDeployInputOutput) DeployTarget() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppDeployInput) string { return v.DeployTarget }).(pulumi.StringOutput)
}

type CreateAppGithubInput struct {
	ProductionBranch string `pulumi:"productionBranch"`
	Url              string `pulumi:"url"`
}

// CreateAppGithubInputInput is an input type that accepts CreateAppGithubInputArgs and CreateAppGithubInputOutput values.
// You can construct a concrete instance of `CreateAppGithubInputInput` via:
//
//	CreateAppGithubInputArgs{...}
type CreateAppGithubInputInput interface {
	pulumi.Input

	ToCreateAppGithubInputOutput() CreateAppGithubInputOutput
	ToCreateAppGithubInputOutputWithContext(context.Context) CreateAppGithubInputOutput
}

type CreateAppGithubInputArgs struct {
	ProductionBranch pulumi.StringInput `pulumi:"productionBranch"`
	Url              pulumi.StringInput `pulumi:"url"`
}

func (CreateAppGithubInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppGithubInput)(nil)).Elem()
}

func (i CreateAppGithubInputArgs) ToCreateAppGithubInputOutput() CreateAppGithubInputOutput {
	return i.ToCreateAppGithubInputOutputWithContext(context.Background())
}

func (i CreateAppGithubInputArgs) ToCreateAppGithubInputOutputWithContext(ctx context.Context) CreateAppGithubInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppGithubInputOutput)
}

func (i CreateAppGithubInputArgs) ToCreateAppGithubInputPtrOutput() CreateAppGithubInputPtrOutput {
	return i.ToCreateAppGithubInputPtrOutputWithContext(context.Background())
}

func (i CreateAppGithubInputArgs) ToCreateAppGithubInputPtrOutputWithContext(ctx context.Context) CreateAppGithubInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppGithubInputOutput).ToCreateAppGithubInputPtrOutputWithContext(ctx)
}

// CreateAppGithubInputPtrInput is an input type that accepts CreateAppGithubInputArgs, CreateAppGithubInputPtr and CreateAppGithubInputPtrOutput values.
// You can construct a concrete instance of `CreateAppGithubInputPtrInput` via:
//
//	        CreateAppGithubInputArgs{...}
//
//	or:
//
//	        nil
type CreateAppGithubInputPtrInput interface {
	pulumi.Input

	ToCreateAppGithubInputPtrOutput() CreateAppGithubInputPtrOutput
	ToCreateAppGithubInputPtrOutputWithContext(context.Context) CreateAppGithubInputPtrOutput
}

type createAppGithubInputPtrType CreateAppGithubInputArgs

func CreateAppGithubInputPtr(v *CreateAppGithubInputArgs) CreateAppGithubInputPtrInput {
	return (*createAppGithubInputPtrType)(v)
}

func (*createAppGithubInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateAppGithubInput)(nil)).Elem()
}

func (i *createAppGithubInputPtrType) ToCreateAppGithubInputPtrOutput() CreateAppGithubInputPtrOutput {
	return i.ToCreateAppGithubInputPtrOutputWithContext(context.Background())
}

func (i *createAppGithubInputPtrType) ToCreateAppGithubInputPtrOutputWithContext(ctx context.Context) CreateAppGithubInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppGithubInputPtrOutput)
}

type CreateAppGithubInputOutput struct{ *pulumi.OutputState }

func (CreateAppGithubInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppGithubInput)(nil)).Elem()
}

func (o CreateAppGithubInputOutput) ToCreateAppGithubInputOutput() CreateAppGithubInputOutput {
	return o
}

func (o CreateAppGithubInputOutput) ToCreateAppGithubInputOutputWithContext(ctx context.Context) CreateAppGithubInputOutput {
	return o
}

func (o CreateAppGithubInputOutput) ToCreateAppGithubInputPtrOutput() CreateAppGithubInputPtrOutput {
	return o.ToCreateAppGithubInputPtrOutputWithContext(context.Background())
}

func (o CreateAppGithubInputOutput) ToCreateAppGithubInputPtrOutputWithContext(ctx context.Context) CreateAppGithubInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateAppGithubInput) *CreateAppGithubInput {
		return &v
	}).(CreateAppGithubInputPtrOutput)
}

func (o CreateAppGithubInputOutput) ProductionBranch() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppGithubInput) string { return v.ProductionBranch }).(pulumi.StringOutput)
}

func (o CreateAppGithubInputOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppGithubInput) string { return v.Url }).(pulumi.StringOutput)
}

type CreateAppGithubInputPtrOutput struct{ *pulumi.OutputState }

func (CreateAppGithubInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateAppGithubInput)(nil)).Elem()
}

func (o CreateAppGithubInputPtrOutput) ToCreateAppGithubInputPtrOutput() CreateAppGithubInputPtrOutput {
	return o
}

func (o CreateAppGithubInputPtrOutput) ToCreateAppGithubInputPtrOutputWithContext(ctx context.Context) CreateAppGithubInputPtrOutput {
	return o
}

func (o CreateAppGithubInputPtrOutput) Elem() CreateAppGithubInputOutput {
	return o.ApplyT(func(v *CreateAppGithubInput) CreateAppGithubInput {
		if v != nil {
			return *v
		}
		var ret CreateAppGithubInput
		return ret
	}).(CreateAppGithubInputOutput)
}

func (o CreateAppGithubInputPtrOutput) ProductionBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateAppGithubInput) *string {
		if v == nil {
			return nil
		}
		return &v.ProductionBranch
	}).(pulumi.StringPtrOutput)
}

func (o CreateAppGithubInputPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateAppGithubInput) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

type CreateAppResourcesInput struct {
	Cpu    string `pulumi:"cpu"`
	Memory string `pulumi:"memory"`
}

// CreateAppResourcesInputInput is an input type that accepts CreateAppResourcesInputArgs and CreateAppResourcesInputOutput values.
// You can construct a concrete instance of `CreateAppResourcesInputInput` via:
//
//	CreateAppResourcesInputArgs{...}
type CreateAppResourcesInputInput interface {
	pulumi.Input

	ToCreateAppResourcesInputOutput() CreateAppResourcesInputOutput
	ToCreateAppResourcesInputOutputWithContext(context.Context) CreateAppResourcesInputOutput
}

type CreateAppResourcesInputArgs struct {
	Cpu    pulumi.StringInput `pulumi:"cpu"`
	Memory pulumi.StringInput `pulumi:"memory"`
}

func (CreateAppResourcesInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppResourcesInput)(nil)).Elem()
}

func (i CreateAppResourcesInputArgs) ToCreateAppResourcesInputOutput() CreateAppResourcesInputOutput {
	return i.ToCreateAppResourcesInputOutputWithContext(context.Background())
}

func (i CreateAppResourcesInputArgs) ToCreateAppResourcesInputOutputWithContext(ctx context.Context) CreateAppResourcesInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateAppResourcesInputOutput)
}

type CreateAppResourcesInputOutput struct{ *pulumi.OutputState }

func (CreateAppResourcesInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateAppResourcesInput)(nil)).Elem()
}

func (o CreateAppResourcesInputOutput) ToCreateAppResourcesInputOutput() CreateAppResourcesInputOutput {
	return o
}

func (o CreateAppResourcesInputOutput) ToCreateAppResourcesInputOutputWithContext(ctx context.Context) CreateAppResourcesInputOutput {
	return o
}

func (o CreateAppResourcesInputOutput) Cpu() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppResourcesInput) string { return v.Cpu }).(pulumi.StringOutput)
}

func (o CreateAppResourcesInputOutput) Memory() pulumi.StringOutput {
	return o.ApplyT(func(v CreateAppResourcesInput) string { return v.Memory }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CreateAppBuildInputInput)(nil)).Elem(), CreateAppBuildInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateAppDeployInputInput)(nil)).Elem(), CreateAppDeployInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateAppGithubInputInput)(nil)).Elem(), CreateAppGithubInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateAppGithubInputPtrInput)(nil)).Elem(), CreateAppGithubInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateAppResourcesInputInput)(nil)).Elem(), CreateAppResourcesInputArgs{})
	pulumi.RegisterOutputType(CreateAppBuildInputOutput{})
	pulumi.RegisterOutputType(CreateAppDeployInputOutput{})
	pulumi.RegisterOutputType(CreateAppGithubInputOutput{})
	pulumi.RegisterOutputType(CreateAppGithubInputPtrOutput{})
	pulumi.RegisterOutputType(CreateAppResourcesInputOutput{})
}

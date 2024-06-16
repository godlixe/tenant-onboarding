package deployer

import (
	"context"
	"fmt"
	"tenant-onboarding/pkg/deployer/types"
)

type Deployer interface {
	GetData(
		ctx context.Context,
		tenantDeploymentJob types.TenantDeploymentJob,
	) (*types.DeploymentSchema, error)

	Initiate(
		ctx context.Context,
		deploymentSchema *types.DeploymentSchema,
	) (*types.RawInfrastructure, error)

	Deploy(
		ctx context.Context,
		deploymentSchema *types.DeploymentSchema,
		rawInfrastructure *types.RawInfrastructure,
	) (*types.RawInfrastructure, error)

	PostDeployment(
		ctx context.Context,
		tenantDeploymentJob types.TenantDeploymentJob,
		rawInfrastructure *types.RawInfrastructure,
		deploymentSchema *types.DeploymentSchema,
	) error

	Cleanup(
		infrastructureDirPath string,
	) error
}

type DeploymentPipeline struct {
	deployer Deployer
}

func NewDeploymentPipeline(
	deployer Deployer,
) *DeploymentPipeline {
	return &DeploymentPipeline{
		deployer: deployer,
	}
}

func (d *DeploymentPipeline) Start(ctx context.Context, tenantDeploymentJob types.TenantDeploymentJob) error {
	deploymentSchema, err := d.deployer.GetData(ctx, tenantDeploymentJob)
	if err != nil {
		return err
	}

	rawInfrastructure, err := d.deployer.Initiate(ctx, deploymentSchema)
	if err != nil {
		return err
	}

	rawInfrastructure, err = d.deployer.Deploy(ctx, deploymentSchema, rawInfrastructure)
	if err != nil {
		return err
	}

	err = d.deployer.PostDeployment(
		ctx,
		tenantDeploymentJob,
		rawInfrastructure,
		deploymentSchema,
	)
	if err != nil {
		return err
	}

	deploymentDir := fmt.Sprintf("/tmp/%v", rawInfrastructure.ID)
	err = d.deployer.Cleanup(deploymentDir)
	if err != nil {
		return err
	}

	return err
}

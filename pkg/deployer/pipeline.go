package deployer

import (
	"context"
	"fmt"
	"tenant-onboarding/pkg/deployer/types"
)

// Deployer defines the basic methods needed
// by the deployment pipeline to run.
type Deployer interface {
	// GetData gathers data needed to deploy infrastructure.
	GetData(
		ctx context.Context,
		tenantDeploymentJob types.TenantDeploymentJob,
	) (*types.DeploymentSchema, error)

	// Initiate prepares the deployment environment and object.
	Initiate(
		ctx context.Context,
		deploymentSchema *types.DeploymentSchema,
	) (*types.RawInfrastructure, error)

	// Deploy runs the deployment.
	Deploy(
		ctx context.Context,
		deploymentSchema *types.DeploymentSchema,
		rawInfrastructure *types.RawInfrastructure,
	) (*types.RawInfrastructure, error)

	// PostDeployment runs everything that is need to be run
	// after deployment, eg: database insertion, event publishing.
	PostDeployment(
		ctx context.Context,
		tenantDeploymentJob types.TenantDeploymentJob,
		rawInfrastructure *types.RawInfrastructure,
		deploymentSchema *types.DeploymentSchema,
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

// Start() initiates the deployment pipeline.
func (d *DeploymentPipeline) Start(ctx context.Context, tenantDeploymentJob types.TenantDeploymentJob) error {
	deploymentSchema, err := d.deployer.GetData(ctx, tenantDeploymentJob)
	if err != nil {
		return err
	}

	fmt.Println("pipeline")

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

	return err
}

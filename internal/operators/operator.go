package operators

import (
	"context"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

type ValidationReply struct {
	IsValid bool
	Reasons []string
}

type ManifestReply struct {
	Manifests map[string]string
}

type API interface {
	ValidateCluster(ctx context.Context, cluster *common.Cluster) (*ValidationReply, error)
	ValidateHost(context.Context, *common.Cluster, *models.Host) (*ValidationReply, error)
	GetCPURequirement(context.Context, *common.Cluster) (int, error)
	GetMemoryRequirement(ctx context.Context, cluster *common.Cluster) (int, error)
	GenerateManifest(context.Context, *common.Cluster) (*ManifestReply, error)
	//TODO: how to handle progress report during the installation?
}

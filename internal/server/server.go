package server

import (
	"github.com/omalloc/contrib/kratos/health"
	"github.com/omalloc/contrib/kratos/registry"
	"github.com/omalloc/contrib/kratos/registry/adapter"
	"github.com/omalloc/contrib/protobuf"
	"go.uber.org/fx"

	"github.com/omalloc/kratos-layout/internal/conf"
	"github.com/omalloc/kratos-layout/internal/data"
	"github.com/omalloc/kratos-layout/internal/service"
)

// ProviderSet is server providers.
var ProviderSet = fx.Provide(
	NewConfigAdapter,
	NewGRPCServer,
	NewHTTPServer,

	adapter.NewEtcdAdapter,
	registry.NewRegistrar,
	registry.NewDiscovery,

	// health..
	health.AsHealth(health.NewServer),
	health.AsChecker[*data.Data](),
	health.AsChecker[*service.GreeterService](),
)

func NewConfigAdapter(bc *conf.Bootstrap) (*protobuf.Registry, *protobuf.Tracing, *conf.Server, *conf.Data) {
	return bc.Registry, bc.Tracing, bc.Server, bc.Data
}

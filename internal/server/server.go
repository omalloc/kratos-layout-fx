package server

import (
	"context"
	"errors"

	"github.com/omalloc/contrib/kratos/health"
	"github.com/omalloc/contrib/kratos/registry"
	"github.com/omalloc/contrib/protobuf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"

	"github.com/omalloc/kratos-layout/internal/conf"
	"github.com/omalloc/kratos-layout/internal/data"
	"github.com/omalloc/kratos-layout/internal/service"
	"github.com/omalloc/kratos-layout/pkg/di"
)

// ProviderSet is server providers.
var ProviderSet = fx.Provide(
	NewConfigAdapter,
	NewEtcdAdapter,

	NewGRPCServer,
	NewHTTPServer,

	registry.NewRegistrar,
	registry.NewDiscovery,

	// health..
	di.AsHealth(health.NewServer),
	di.AsChecker[*data.Data](),
	di.AsChecker[*service.GreeterService](),
)

func NewConfigAdapter(bc *conf.Bootstrap) (*protobuf.Registry, *protobuf.Tracing, *conf.Server, *conf.Data) {
	return bc.Registry, bc.Tracing, bc.Server, bc.Data
}

func NewEtcdAdapter(lc fx.Lifecycle, c *protobuf.Registry) (*clientv3.Client, error) {
	cli, cleanup, err := registry.NewEtcd(c)
	if err != nil {
		return nil, errors.New("failed to connect to etcd: " + err.Error())
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cleanup()
			return nil
		},
	})
	return cli, nil
}

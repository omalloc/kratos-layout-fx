package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/omalloc/contrib/kratos/health"
	"go.uber.org/fx"

	"github.com/omalloc/kratos-layout/internal/biz"
	"github.com/omalloc/kratos-layout/internal/conf"
	"github.com/omalloc/kratos-layout/internal/data"
	"github.com/omalloc/kratos-layout/internal/server"
	"github.com/omalloc/kratos-layout/internal/service"
)

func initApp(logger log.Logger, c config.Config, bc *conf.Bootstrap) *fx.App {
	return fx.New(
		fx.NopLogger,

		AsAnno(c, new(config.Config)),
		AsAnno(logger, new(log.Logger)),
		As(bc),

		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,

		fx.Invoke(newApp),
	)
}

func newApp(lc fx.Lifecycle, logger log.Logger, registrar registry.Registrar, gs *grpc.Server, hs *http.Server, hh *health.Server) {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(registrar),
		kratos.Server(
			gs,
			hs,
			hh,
		),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Run(); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Stop()
		},
	})
}

func AsAnno(t interface{}, interfaces ...interface{}) fx.Option {
	return fx.Supply(fx.Annotate(t, fx.As(interfaces...)))
}

func As(values interface{}) fx.Option {
	return fx.Supply(values)
}

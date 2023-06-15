package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/fx"

	"github.com/omalloc/kratos-layout/internal/conf"
	"github.com/omalloc/kratos-layout/pkg/di"
)

// ProviderSet is data providers.
var ProviderSet = fx.Options(
	fx.Provide(
		NewDataAdapter,
		NewGreeterRepo,
	),

	// health checker..
	di.AsHealthChecker[*Data](),
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

// Check data health checker, implement health.Checker interface.
func (d *Data) Check(ctx context.Context) error {
	return nil
}

func NewDataAdapter(lc fx.Lifecycle, c *conf.Data, logger log.Logger) (*Data, error) {
	data, cleanup, err := NewData(c, logger)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cleanup()
			return nil
		},
	})
	return data, nil
}

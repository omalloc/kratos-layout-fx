package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/kratos/orm"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/omalloc/kratos-layout/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = fx.Provide(
	NewDataAdapter,
	NewGreeterRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	db, err := orm.New(
		orm.WithDriver(getDataDriver(c.Database.Driver, c.Database.Source)),
		orm.WithTracing(),
		orm.WithLogger(
			orm.WithLogHelper(logger),
			orm.WIthSlowThreshold(time.Second*2),
			orm.WithDebug(),
		),
	)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		if db != nil {
			if sql, err := db.DB(); err == nil {
				sql.Close()
			}
		}
	}

	return &Data{
		db: db,
	}, cleanup, nil
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

func getDataDriver(typo, source string) gorm.Dialector {
	switch typo {
	case "sqlite":
		return sqlite.Open(source)
	default:
		return mysql.Open(source)
	}
}

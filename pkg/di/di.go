package di

import (
	"github.com/omalloc/contrib/kratos/health"
	"go.uber.org/fx"
)

func As[T health.Checker](f any) {
}

func AsHealthChecker[T health.Checker]() fx.Option {
	return fx.Provide(AsChecker[T]())
}

func AsChecker[T health.Checker]() interface{} {
	return fx.Annotate(
		func(typo T) health.Checker {
			return typo
		},
		fx.ResultTags(`group:"health"`),
	)
}

func AsHealth(f any) any {
	return fx.Annotate(
		f,
		fx.ParamTags(`group:"health"`),
	)
}

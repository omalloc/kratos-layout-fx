package service

import (
	"go.uber.org/fx"

	"github.com/omalloc/kratos-layout/pkg/di"
)

// ProviderSet is service providers.
var ProviderSet = fx.Provide(
	// all services..
	NewGreeterService,

	// all checkers..
	di.AsChecker[*GreeterService](),
)

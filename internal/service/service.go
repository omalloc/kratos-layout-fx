package service

import (
	"go.uber.org/fx"
)

// ProviderSet is service providers.
var ProviderSet = fx.Provide(
	// all services..
	NewGreeterService,
)

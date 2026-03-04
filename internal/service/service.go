package service

import "github.com/google/wire"

// ProviderSet is service providers. test
var ProviderSet = wire.NewSet(NewHealthService)

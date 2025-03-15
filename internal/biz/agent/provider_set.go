package agent

import (
	"github.com/google/wire"
)

// ProviderSet is wire providers.
var ProviderSet = wire.NewSet(NewHub)

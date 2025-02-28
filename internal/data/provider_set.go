package data

import (
	"github.com/google/wire"

	"github.com/iter-x/iter-x/internal/data/impl"
)

var ProviderSet = wire.NewSet(
	NewConnection,
	NewTx,
	NewTransactionRepository,
	impl.NewAuth,
	impl.NewTrip,
	impl.NewPointsOfInterest,
)

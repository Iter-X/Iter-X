package impl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConnection,
	NewAuthRepository,
	NewTripRepository,
	NewTransactionRepository,
	NewRefreshTokenRepository,
	NewTripDailyRepository,
	NewTripDailyItemRepository,
)

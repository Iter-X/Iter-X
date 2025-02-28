package impl

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewAuth,
	NewRefreshToken,
	NewTrip,
	NewDailyTrip,
	NewDailyItinerary,
	NewPointsOfInterest,
	NewCity,
	NewContinent,
	NewCountry,
	NewState,
)

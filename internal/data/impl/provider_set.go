package impl

import (
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewAgent,
	NewAuth,
	NewUser,
	NewLanguage,
	NewRefreshToken,
	NewTrip,
	NewDailyTrip,
	NewDailyItinerary,
	NewPointsOfInterest,
	NewCity,
	NewContinent,
	NewCountry,
	NewState,
	NewFiles,
	NewPoiFiles,
	NewTool,
)

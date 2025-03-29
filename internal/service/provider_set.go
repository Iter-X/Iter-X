package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAuth,
	NewUser,
	NewTrip,
	NewPointsOfInterest,
	NewGeo,
	NewStorage,
)

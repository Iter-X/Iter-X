package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuth, NewTrip, NewPointsOfInterestService, NewGeoService)

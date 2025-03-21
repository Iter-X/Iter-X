package service

import (
	"context"

	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/service/build"
)

// GeoService geographic information service
type GeoService struct {
	geoV1.UnimplementedGeoServiceServer
	geoBiz *biz.Geo
}

// NewGeoService creates a new geographic information service instance
func NewGeoService(geoBiz *biz.Geo) *GeoService {
	return &GeoService{
		geoBiz: geoBiz,
	}
}

// ListContinents lists all continents
func (s *GeoService) ListContinents(ctx context.Context, req *geoV1.ListContinentsRequest) (*geoV1.ListContinentsResponse, error) {
	// Convert PB to BO
	params := &bo.ListContinentsParams{
		Limit:  req.Limit,
		Offset: int(req.Offset),
	}

	// Call biz layer
	continents, total, err := s.geoBiz.ListContinents(ctx, params)
	if err != nil {
		return nil, err
	}

	// Convert result to PB response
	resp := &geoV1.ListContinentsResponse{
		Total:      total,
		Continents: build.ToContinentsProto(continents),
	}

	return resp, nil
}

// ListCountries lists countries, optionally filtered by continent
func (s *GeoService) ListCountries(ctx context.Context, req *geoV1.ListCountriesRequest) (*geoV1.ListCountriesResponse, error) {
	// Convert PB to BO
	params := &bo.ListCountriesParams{
		ContinentID: uint(req.ContinentId),
		Limit:       req.Limit,
		Offset:      int(req.Offset),
	}

	// Call biz layer
	countries, total, err := s.geoBiz.ListCountries(ctx, params)
	if err != nil {
		return nil, err
	}

	// Convert result to PB response
	resp := &geoV1.ListCountriesResponse{
		Total:     total,
		Countries: build.ToCountriesProto(countries),
	}

	return resp, nil
}

// ListStates lists states/provinces, optionally filtered by country
func (s *GeoService) ListStates(ctx context.Context, req *geoV1.ListStatesRequest) (*geoV1.ListStatesResponse, error) {
	// Convert PB to BO
	params := &bo.ListStatesParams{
		CountryID: uint(req.CountryId),
		Limit:     req.Limit,
		Offset:    int(req.Offset),
	}

	// Call biz layer
	states, total, err := s.geoBiz.ListStates(ctx, params)
	if err != nil {
		return nil, err
	}

	// Convert result to PB response
	resp := &geoV1.ListStatesResponse{
		Total:  total,
		States: build.ToStatesProto(states),
	}

	return resp, nil
}

// ListCities lists cities, optionally filtered by state/province
func (s *GeoService) ListCities(ctx context.Context, req *geoV1.ListCitiesRequest) (*geoV1.ListCitiesResponse, error) {
	// Convert PB to BO
	params := &bo.ListCitiesParams{
		StateID: uint(req.StateId),
		Limit:   req.Limit,
		Offset:  int(req.Offset),
	}

	// Call biz layer
	cities, total, err := s.geoBiz.ListCities(ctx, params)
	if err != nil {
		return nil, err
	}

	// Convert result to PB response
	resp := &geoV1.ListCitiesResponse{
		Total:  total,
		Cities: build.ToCitiesProto(cities),
	}

	return resp, nil
}

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

// NewGeo creates a new geographic information service instance
func NewGeo(geoBiz *biz.Geo) *GeoService {
	return &GeoService{
		geoBiz: geoBiz,
	}
}

// ListContinents lists all continents
func (s *GeoService) ListContinents(ctx context.Context, req *geoV1.ListContinentsRequest) (*geoV1.ListContinentsResponse, error) {
	// Convert PB to BO
	params := &bo.ListContinentsParams{}
	if req.Pagination != nil {
		params.Pagination = bo.FromPageAndSize(req.Pagination.Page, req.Pagination.Size)
	}

	// Call biz
	continents, total, err := s.geoBiz.ListContinents(ctx, params)
	if err != nil {
		return nil, err
	}

	return &geoV1.ListContinentsResponse{
		Total:      total,
		Continents: build.ToContinentsProto(continents),
	}, nil
}

// ListCountries lists countries, optionally filtered by continent
func (s *GeoService) ListCountries(ctx context.Context, req *geoV1.ListCountriesRequest) (*geoV1.ListCountriesResponse, error) {
	// Convert PB to BO
	params := &bo.ListCountriesParams{
		ContinentID: uint(req.ContinentId),
	}
	if req.Pagination != nil {
		params.Pagination = bo.FromPageAndSize(req.Pagination.Page, req.Pagination.Size)
	}

	// Call biz
	countries, total, err := s.geoBiz.ListCountries(ctx, params)
	if err != nil {
		return nil, err
	}

	return &geoV1.ListCountriesResponse{
		Total:     total,
		Countries: build.ToCountriesProto(countries),
	}, nil
}

// ListStates lists states/provinces, optionally filtered by country
func (s *GeoService) ListStates(ctx context.Context, req *geoV1.ListStatesRequest) (*geoV1.ListStatesResponse, error) {
	// Convert PB to BO
	params := &bo.ListStatesParams{
		CountryID: uint(req.CountryId),
	}
	if req.Pagination != nil {
		params.Pagination = bo.FromPageAndSize(req.Pagination.Page, req.Pagination.Size)
	}

	// Call biz
	states, total, err := s.geoBiz.ListStates(ctx, params)
	if err != nil {
		return nil, err
	}

	return &geoV1.ListStatesResponse{
		Total:  total,
		States: build.ToStatesProto(states),
	}, nil
}

// ListCities lists cities, optionally filtered by state/province
func (s *GeoService) ListCities(ctx context.Context, req *geoV1.ListCitiesRequest) (*geoV1.ListCitiesResponse, error) {
	// Convert PB to BO
	params := &bo.ListCitiesParams{
		StateID: uint(req.StateId),
	}
	if req.Pagination != nil {
		params.Pagination = bo.FromPageAndSize(req.Pagination.Page, req.Pagination.Size)
	}

	// Call biz
	cities, total, err := s.geoBiz.ListCities(ctx, params)
	if err != nil {
		return nil, err
	}

	return &geoV1.ListCitiesResponse{
		Total:  total,
		Cities: build.ToCitiesProto(cities),
	}, nil
}

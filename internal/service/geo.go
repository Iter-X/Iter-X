package service

import (
	"context"

	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
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
		Total: total,
	}

	for _, continent := range continents {
		resp.Continents = append(resp.Continents, &geoV1.Continent{
			Id:     uint32(continent.ID),
			Name:   continent.Name,
			NameEn: continent.NameEn,
			NameCn: continent.NameCn,
			Code:   continent.Code,
		})
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
		Total: total,
	}

	for _, country := range countries {
		countryProto := &geoV1.Country{
			Id:          uint32(country.ID),
			Name:        country.Name,
			NameEn:      country.NameEn,
			NameCn:      country.NameCn,
			Code:        country.Code,
			ContinentId: uint32(country.ContinentID),
		}

		if country.Continent != nil {
			countryProto.Continent = &geoV1.Continent{
				Id:     uint32(country.Continent.ID),
				Name:   country.Continent.Name,
				NameEn: country.Continent.NameEn,
				NameCn: country.Continent.NameCn,
				Code:   country.Continent.Code,
			}
		}

		resp.Countries = append(resp.Countries, countryProto)
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
		Total: total,
	}

	for _, state := range states {
		stateProto := &geoV1.State{
			Id:        uint32(state.ID),
			Name:      state.Name,
			NameEn:    state.NameEn,
			NameCn:    state.NameCn,
			Code:      state.Code,
			CountryId: uint32(state.CountryID),
		}

		if state.Country != nil {
			stateProto.Country = &geoV1.Country{
				Id:     uint32(state.Country.ID),
				Name:   state.Country.Name,
				NameEn: state.Country.NameEn,
				NameCn: state.Country.NameCn,
				Code:   state.Country.Code,
			}
		}

		resp.States = append(resp.States, stateProto)
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
		Total: total,
	}

	for _, city := range cities {
		cityProto := &geoV1.City{
			Id:      uint32(city.ID),
			Name:    city.Name,
			NameEn:  city.NameEn,
			NameCn:  city.NameCn,
			Code:    city.Code,
			StateId: uint32(city.StateID),
		}

		if city.State != nil {
			cityProto.State = &geoV1.State{
				Id:     uint32(city.State.ID),
				Name:   city.State.Name,
				NameEn: city.State.NameEn,
				NameCn: city.State.NameCn,
				Code:   city.State.Code,
			}
		}

		resp.Cities = append(resp.Cities, cityProto)
	}

	return resp, nil
}

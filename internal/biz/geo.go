package biz

import (
	"context"

	"go.uber.org/zap"

	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/repository"
)

// Geo 地理信息业务逻辑
type Geo struct {
	continentRepo repository.ContinentRepo
	countryRepo   repository.CountryRepo
	stateRepo     repository.StateRepo
	cityRepo      repository.CityRepo
	logger        *zap.SugaredLogger
}

// NewGeo 创建地理信息业务逻辑实例
func NewGeo(
	continentRepo repository.ContinentRepo,
	countryRepo repository.CountryRepo,
	stateRepo repository.StateRepo,
	cityRepo repository.CityRepo,
	logger *zap.SugaredLogger,
) *Geo {
	return &Geo{
		continentRepo: continentRepo,
		countryRepo:   countryRepo,
		stateRepo:     stateRepo,
		cityRepo:      cityRepo,
		logger:        logger.Named("biz.geo"),
	}
}

// ListContinents 列出所有大洲
func (g *Geo) ListContinents(ctx context.Context, req *geoV1.ListContinentsRequest) (*geoV1.ListContinentsResponse, error) {
	params := &bo.ListContinentsParams{
		PageSize:  req.PageSize,
		PageToken: req.PageToken,
	}

	continents, pagination, err := g.continentRepo.ListContinents(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &geoV1.ListContinentsResponse{
		NextPageToken: pagination.NextPageToken,
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

// ListCountries 列出国家，可选按大洲过滤
func (g *Geo) ListCountries(ctx context.Context, req *geoV1.ListCountriesRequest) (*geoV1.ListCountriesResponse, error) {
	params := &bo.ListCountriesParams{
		ContinentID: uint(req.ContinentId),
		PageSize:    req.PageSize,
		PageToken:   req.PageToken,
	}

	countries, pagination, err := g.countryRepo.ListCountries(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &geoV1.ListCountriesResponse{
		NextPageToken: pagination.NextPageToken,
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

// ListStates 列出州/省，可选按国家过滤
func (g *Geo) ListStates(ctx context.Context, req *geoV1.ListStatesRequest) (*geoV1.ListStatesResponse, error) {
	params := &bo.ListStatesParams{
		CountryID: uint(req.CountryId),
		PageSize:  req.PageSize,
		PageToken: req.PageToken,
	}

	states, pagination, err := g.stateRepo.ListStates(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &geoV1.ListStatesResponse{
		NextPageToken: pagination.NextPageToken,
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

// ListCities 列出城市，可选按州/省过滤
func (g *Geo) ListCities(ctx context.Context, req *geoV1.ListCitiesRequest) (*geoV1.ListCitiesResponse, error) {
	params := &bo.ListCitiesParams{
		StateID:   uint(req.StateId),
		PageSize:  req.PageSize,
		PageToken: req.PageToken,
	}

	cities, pagination, err := g.cityRepo.ListCities(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &geoV1.ListCitiesResponse{
		NextPageToken: pagination.NextPageToken,
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

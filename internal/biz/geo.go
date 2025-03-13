package biz

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
)

// Geo geographic information business logic
type Geo struct {
	continentRepo repository.ContinentRepo
	countryRepo   repository.CountryRepo
	stateRepo     repository.StateRepo
	cityRepo      repository.CityRepo
	logger        *zap.SugaredLogger
}

// NewGeo creates a new geographic information business logic instance
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

// ListContinents lists all continents
func (g *Geo) ListContinents(ctx context.Context, params *bo.ListContinentsParams) ([]*do.Continent, int64, error) {
	// Call repository layer
	return g.continentRepo.ListContinents(ctx, params)
}

// ListCountries lists countries, optionally filtered by continent
func (g *Geo) ListCountries(ctx context.Context, params *bo.ListCountriesParams) ([]*do.Country, int64, error) {
	// Call repository layer
	return g.countryRepo.ListCountries(ctx, params)
}

// ListStates lists states/provinces, optionally filtered by country
func (g *Geo) ListStates(ctx context.Context, params *bo.ListStatesParams) ([]*do.State, int64, error) {
	// Call repository layer
	return g.stateRepo.ListStates(ctx, params)
}

// ListCities lists cities, optionally filtered by state/province
func (g *Geo) ListCities(ctx context.Context, params *bo.ListCitiesParams) ([]*do.City, int64, error) {
	// Call repository layer
	return g.cityRepo.ListCities(ctx, params)
}

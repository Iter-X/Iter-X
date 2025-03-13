package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/country"
)

// NewCountry creates a new country repository implementation
func NewCountry(d *data.Data, continentRepository repository.ContinentRepo, logger *zap.SugaredLogger) repository.CountryRepo {
	return &countryRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.country"),
		continentRepository:            continentRepository,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		stateRepositoryImpl:            new(stateRepositoryImpl),
		continentRepositoryImpl:        new(continentRepositoryImpl),
	}
}

type countryRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	continentRepository repository.ContinentRepo

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	stateRepositoryImpl            repository.BaseRepo[*ent.State, *do.State]
	continentRepositoryImpl        repository.BaseRepo[*ent.Continent, *do.Continent]
}

// ToEntity converts a persistent object to a domain object
func (c *countryRepositoryImpl) ToEntity(po *ent.Country) *do.Country {
	if po == nil {
		return nil
	}

	return &do.Country{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		Name:        po.Name,
		NameEn:      po.NameEn,
		NameCn:      po.NameCn,
		Code:        po.Code,
		ContinentID: po.ContinentID,
		Poi:         c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		State:       c.stateRepositoryImpl.ToEntities(po.Edges.State),
		Continent:   c.continentRepositoryImpl.ToEntity(po.Edges.Continent),
	}
}

// ToEntities converts a collection of persistent objects to domain objects
func (c *countryRepositoryImpl) ToEntities(pos []*ent.Country) []*do.Country {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.Country, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}

// SearchPointsOfInterest searches for points of interest
func (c *countryRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsCountry() {
		return c.continentRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := c.GetTx(ctx).Country
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(country.Or(
			country.NameContains(keyword),
			country.NameCnContains(keyword),
			country.NameEnContains(keyword),
			country.CodeContains(keyword),
		)).
		WithContinent().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		countryDo := c.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			Country:   countryDo,
			Continent: countryDo.Continent,
		})
	}
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := c.continentRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}

	return pois, nil
}

// ListCountries lists countries, optionally filtered by continent
func (c *countryRepositoryImpl) ListCountries(ctx context.Context, params *bo.ListCountriesParams) ([]*do.Country, int64, error) {
	query := c.GetTx(ctx).Country.Query()

	// Filter by continent if specified
	if params.ContinentID > 0 {
		query = query.Where(country.ContinentID(params.ContinentID))
	}

	// Set pagination
	limit := int(params.Limit)
	if limit <= 0 {
		limit = 10 // Default to 10 records per page
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	query = query.Offset(params.Offset).Limit(limit)

	// Load related continent information
	query = query.WithContinent()

	// Execute query
	countries, err := query.Order(ent.Asc(country.FieldName)).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Convert to domain objects
	result := make([]*do.Country, len(countries))
	for i, country := range countries {
		result[i] = c.ToEntity(country)
	}

	return result, int64(total), nil
}

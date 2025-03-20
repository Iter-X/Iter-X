package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/continent"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

// NewContinent creates a new continent repository implementation
func NewContinent(d *data.Data, logger *zap.SugaredLogger) repository.ContinentRepo {
	return &continentRepositoryImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.continent"),
	}
}

type continentRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

// ToEntity converts a persistent object to a domain object
func (c *continentRepositoryImpl) ToEntity(po *ent.Continent) *do.Continent {
	if po == nil {
		return nil
	}
	return build.ContinentRepositoryImplToEntity(po)
}

// ToEntities converts a collection of persistent objects to domain objects
func (c *continentRepositoryImpl) ToEntities(pos []*ent.Continent) []*do.Continent {
	if pos == nil {
		return nil
	}
	return build.ContinentRepositoryImplToEntities(pos)
}

// SearchPointsOfInterest searches for points of interest
func (c *continentRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	cli := c.GetTx(ctx).Continent
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(continent.Or(
			continent.NameContains(keyword),
			continent.NameCnContains(keyword),
			continent.NameEnContains(keyword),
			continent.CodeContains(keyword),
		)).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		continentDo := c.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			Continent: continentDo,
		})
	}
	return pois, nil
}

// ListContinents lists all continents
func (c *continentRepositoryImpl) ListContinents(ctx context.Context, params *bo.ListContinentsParams) ([]*do.Continent, int64, error) {
	query := c.GetTx(ctx).Continent.Query()

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

	// Execute query
	continents, err := query.Order(ent.Asc(continent.FieldName)).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Convert to domain objects
	result := make([]*do.Continent, len(continents))
	for i, continent := range continents {
		result[i] = c.ToEntity(continent)
	}

	return result, int64(total), nil
}

package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/data/ent/continent"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewContinent(d *data.Data, logger *zap.SugaredLogger) repository.ContinentRepo {
	return &continentRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.continent"),
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		countryRepositoryImpl:          new(countryRepositoryImpl),
	}
}

type continentRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	countryRepositoryImpl          repository.BaseRepo[*ent.Country, *do.Country]
}

func (c *continentRepositoryImpl) ToEntity(po *ent.Continent) *do.Continent {
	if po == nil {
		return nil
	}
	return &do.Continent{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		Poi:       c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		Country:   c.countryRepositoryImpl.ToEntities(po.Edges.Country),
	}
}

func (c *continentRepositoryImpl) ToEntities(pos []*ent.Continent) []*do.Continent {
	if pos == nil {
		return nil
	}
	list := make([]*do.Continent, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}

func (c *continentRepositoryImpl) SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error) {
	cli := c.GetTx(ctx).Continent

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

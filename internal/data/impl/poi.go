package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/pointsofinterest"
)

func NewPointsOfInterest(d *data.Data, cityRepository repository.CityRepo, logger *zap.SugaredLogger) repository.PointsOfInterestRepo {
	return &pointsOfInterestRepositoryImpl{
		Tx:                           d.Tx,
		logger:                       logger.Named("repo.poi"),
		cityRepository:               cityRepository,
		cityRepositoryImpl:           new(cityRepositoryImpl),
		stateRepositoryImpl:          new(stateRepositoryImpl),
		countryRepositoryImpl:        new(countryRepositoryImpl),
		continentRepositoryImpl:      new(continentRepositoryImpl),
		dailyItineraryRepositoryImpl: new(dailyItineraryRepositoryImpl),
	}
}

type pointsOfInterestRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	cityRepository repository.CityRepo

	cityRepositoryImpl           repository.BaseRepo[*ent.City, *do.City]
	stateRepositoryImpl          repository.BaseRepo[*ent.State, *do.State]
	countryRepositoryImpl        repository.BaseRepo[*ent.Country, *do.Country]
	continentRepositoryImpl      repository.BaseRepo[*ent.Continent, *do.Continent]
	dailyItineraryRepositoryImpl repository.BaseRepo[*ent.DailyItinerary, *do.DailyItinerary]
}

func (r *pointsOfInterestRepositoryImpl) ToEntity(po *ent.PointsOfInterest) *do.PointsOfInterest {
	if po == nil {
		return nil
	}
	return &do.PointsOfInterest{
		ID:                         po.ID,
		CreatedAt:                  po.CreatedAt,
		UpdatedAt:                  po.UpdatedAt,
		Name:                       po.Name,
		NameEn:                     po.NameEn,
		NameCn:                     po.NameCn,
		Description:                po.Description,
		Address:                    po.Address,
		Latitude:                   po.Latitude,
		Longitude:                  po.Longitude,
		Type:                       po.Type,
		Category:                   po.Category,
		Rating:                     po.Rating,
		RecommendedDurationMinutes: po.RecommendedDurationMinutes,
		CityID:                     po.CityID,
		StateID:                    po.StateID,
		CountryID:                  po.CountryID,
		ContinentID:                po.ContinentID,
		City:                       r.cityRepositoryImpl.ToEntity(po.Edges.City),
		State:                      r.stateRepositoryImpl.ToEntity(po.Edges.State),
		Country:                    r.countryRepositoryImpl.ToEntity(po.Edges.Country),
		Continent:                  r.continentRepositoryImpl.ToEntity(po.Edges.Continent),
		DailyItinerary:             r.dailyItineraryRepositoryImpl.ToEntities(po.Edges.DailyItinerary),
	}
}

func (r *pointsOfInterestRepositoryImpl) ToEntities(pos []*ent.PointsOfInterest) []*do.PointsOfInterest {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.PointsOfInterest, 0, len(pos))
	for _, v := range pos {
		list = append(list, r.ToEntity(v))
	}
	return list
}

func (r *pointsOfInterestRepositoryImpl) SearchPointsOfInterest(ctx context.Context, keyword string, limit int) ([]*do.PointsOfInterest, error) {
	cli := r.GetTx(ctx).PointsOfInterest

	rows, err := cli.Query().
		Where(pointsofinterest.Or(
			pointsofinterest.NameContains(keyword),
			pointsofinterest.NameCnContains(keyword),
			pointsofinterest.NameEnContains(keyword),
			pointsofinterest.DescriptionContains(keyword),
		)).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return r.cityRepository.SearchPointsOfInterest(ctx, keyword, limit)
	}
	return r.ToEntities(rows), nil
}

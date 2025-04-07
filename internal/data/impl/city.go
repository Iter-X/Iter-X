package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/city"
	"github.com/iter-x/iter-x/internal/data/ent/state"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewCity(d *data.Data, stateRepository repository.StateRepo, logger *zap.SugaredLogger) repository.CityRepo {
	return &cityRepositoryImpl{
		Tx:              d.Tx,
		logger:          logger.Named("repo.city"),
		stateRepository: stateRepository,
	}
}

type cityRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	stateRepository repository.StateRepo
}

func (r *cityRepositoryImpl) ToEntity(po *ent.City) *do.City {
	if po == nil {
		return nil
	}
	return build.CityRepositoryImplToEntity(po)
}

func (r *cityRepositoryImpl) ToEntities(pos []*ent.City) []*do.City {
	if len(pos) == 0 {
		return nil
	}
	return build.CityRepositoryImplToEntities(pos)
}

func (r *cityRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsCity() {
		return r.stateRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := r.GetTx(ctx).City
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(city.Or(
			city.NameLocalContains(keyword),
			city.NameCnContains(keyword),
			city.NameEnContains(keyword),
			city.CodeContains(keyword),
		)).
		WithState(func(stateQuery *ent.StateQuery) {
			stateQuery.WithCountry(func(countryQuery *ent.CountryQuery) {
				countryQuery.WithContinent()
			})
		}).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		cityDo := r.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			City:      cityDo,
			State:     cityDo.State,
			Country:   cityDo.State.Country,
			Continent: cityDo.State.Country.Continent,
		})
	}
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := r.stateRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}

	return pois, nil
}

// ListCities lists cities, optionally filtered by state/province
func (r *cityRepositoryImpl) ListCities(ctx context.Context, params *bo.ListCitiesParams) ([]*do.City, int64, error) {
	query := r.GetTx(ctx).City.Query()

	// Filter by state if specified
	if params.StateId != nil {
		query = query.Where(city.StateID(uint(*params.StateId)))
	}

	// Filter by country if specified
	if params.CountryId != nil {
		query = query.Where(city.HasStateWith(state.CountryID(uint(*params.CountryId))))
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	query = query.Offset(params.GetOffset4Db()).Limit(params.GetLimit4Db())

	// Load related state information
	query = query.WithState()

	// Execute query
	cities, err := query.Order(ent.Asc(city.FieldNameEn)).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Convert to DO objects
	result := make([]*do.City, len(cities))
	for i, c := range cities {
		result[i] = r.ToEntity(c)
	}

	return result, int64(total), nil
}

func (r *cityRepositoryImpl) GetCityIdByName(ctx context.Context, cityName string) (int32, error) {
	cli := r.GetTx(ctx).City

	// 先尝试精确匹配本地名称
	row, err := cli.Query().
		Where(city.NameLocalEQ(cityName)).
		First(ctx)
	if err != nil {
		// 如果精确匹配失败，尝试模糊匹配
		row, err = cli.Query().
			Where(city.NameLocalContains(cityName)).
			Order(ent.Asc(city.FieldNameLocal)). // 按名称排序，确保结果一致性
			First(ctx)
		if err != nil {
			// 如果本地名称都匹配失败，尝试英文名称
			row, err = cli.Query().
				Where(city.NameEnContains(cityName)).
				Order(ent.Asc(city.FieldNameEn)).
				First(ctx)
			if err != nil {
				return 0, err
			}
		}
	}

	return int32(row.ID), nil
}

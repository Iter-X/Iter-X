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
)

func NewCity(d *data.Data, stateRepository repository.StateRepo, logger *zap.SugaredLogger) repository.CityRepo {
	return &cityRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.city"),
		stateRepository:                stateRepository,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		stateRepositoryImpl:            new(stateRepositoryImpl),
	}
}

type cityRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	stateRepository repository.StateRepo

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	stateRepositoryImpl            repository.BaseRepo[*ent.State, *do.State]
}

func (c *cityRepositoryImpl) ToEntity(po *ent.City) *do.City {
	if po == nil {
		return nil
	}
	return &do.City{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		StateID:   po.StateID,
		Poi:       c.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		State:     c.stateRepositoryImpl.ToEntity(po.Edges.State),
	}
}

func (c *cityRepositoryImpl) ToEntities(pos []*ent.City) []*do.City {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.City, 0, len(pos))
	for _, v := range pos {
		list = append(list, c.ToEntity(v))
	}
	return list
}

func (c *cityRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsCity() {
		return c.stateRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := c.GetTx(ctx).City
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(city.Or(
			city.NameContains(keyword),
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
		cityDo := c.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			City:      cityDo,
			State:     cityDo.State,
			Country:   cityDo.State.Country,
			Continent: cityDo.State.Country.Continent,
		})
	}
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := c.stateRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}

	return pois, nil
}

// ListCities 列出城市，可选按州/省过滤
func (r *cityRepositoryImpl) ListCities(ctx context.Context, params *bo.ListCitiesParams) ([]*do.City, *bo.PaginationResult, error) {
	query := r.GetTx(ctx).City.Query()

	// 按州/省过滤
	if params.StateID > 0 {
		query = query.Where(city.StateID(params.StateID))
	}

	// 解析分页令牌
	var err error
	if params.Offset == 0 && params.PageToken != "" {
		params.Offset, err = bo.ParsePageToken(params.PageToken)
		if err != nil {
			return nil, nil, err
		}
	}

	// 设置分页
	limit := int(params.PageSize)
	if limit <= 0 {
		limit = 10 // 默认每页10条
	}

	query = query.Offset(params.Offset).Limit(limit + 1) // 多查询一条用于判断是否有更多数据

	// 加载关联的州/省信息
	query = query.WithState()

	// 执行查询
	cities, err := query.Order(ent.Asc(city.FieldName)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 判断是否有更多数据
	hasMore := false
	if len(cities) > limit {
		hasMore = true
		cities = cities[:limit] // 去掉多查询的一条
	}

	// 计算下一页的偏移量
	nextOffset := params.Offset + len(cities)

	// 转换为DO对象
	result := make([]*do.City, len(cities))
	for i, c := range cities {
		result[i] = r.ToEntity(c)
	}

	// 生成下一页令牌
	nextPageToken := bo.GenerateNextPageToken(nextOffset, hasMore)

	return result, &bo.PaginationResult{
		NextPageToken: nextPageToken,
		HasMore:       hasMore,
	}, nil
}

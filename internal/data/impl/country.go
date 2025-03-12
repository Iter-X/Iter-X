package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/data/ent/country"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

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

// ListCountries 列出国家，可选按大洲过滤
func (r *countryRepositoryImpl) ListCountries(ctx context.Context, params *bo.ListCountriesParams) ([]*do.Country, *bo.PaginationResult, error) {
	query := r.GetTx(ctx).Country.Query()

	// 按大洲过滤
	if params.ContinentID > 0 {
		query = query.Where(country.ContinentID(params.ContinentID))
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

	// 加载关联的大洲信息
	query = query.WithContinent()

	// 执行查询
	countries, err := query.Order(ent.Asc(country.FieldName)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 判断是否有更多数据
	hasMore := false
	if len(countries) > limit {
		hasMore = true
		countries = countries[:limit] // 去掉多查询的一条
	}

	// 计算下一页的偏移量
	nextOffset := params.Offset + len(countries)

	// 转换为DO对象
	result := make([]*do.Country, len(countries))
	for i, c := range countries {
		result[i] = r.ToEntity(c)
	}

	// 生成下一页令牌
	nextPageToken := bo.GenerateNextPageToken(nextOffset, hasMore)

	return result, &bo.PaginationResult{
		NextPageToken: nextPageToken,
		HasMore:       hasMore,
	}, nil
}

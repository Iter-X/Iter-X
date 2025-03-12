package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
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

// ListContinents 列出所有大洲
func (r *continentRepositoryImpl) ListContinents(ctx context.Context, params *bo.ListContinentsParams) ([]*do.Continent, *bo.PaginationResult, error) {
	query := r.GetTx(ctx).Continent.Query()

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

	// 执行查询
	continents, err := query.Order(ent.Asc(continent.FieldName)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 判断是否有更多数据
	hasMore := false
	if len(continents) > limit {
		hasMore = true
		continents = continents[:limit] // 去掉多查询的一条
	}

	// 计算下一页的偏移量
	nextOffset := params.Offset + len(continents)

	// 转换为DO对象
	result := make([]*do.Continent, len(continents))
	for i, c := range continents {
		result[i] = r.ToEntity(c)
	}

	// 生成下一页令牌
	nextPageToken := bo.GenerateNextPageToken(nextOffset, hasMore)

	return result, &bo.PaginationResult{
		NextPageToken: nextPageToken,
		HasMore:       hasMore,
	}, nil
}

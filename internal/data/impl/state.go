package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/data/ent/state"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewState(d *data.Data, countryRepository repository.CountryRepo, logger *zap.SugaredLogger) repository.StateRepo {
	return &stateRepositoryImpl{
		Tx:                             d.Tx,
		logger:                         logger.Named("repo.state"),
		countryRepository:              countryRepository,
		pointsOfInterestRepositoryImpl: new(pointsOfInterestRepositoryImpl),
		cityRepositoryImpl:             new(cityRepositoryImpl),
		countryRepositoryImpl:          new(countryRepositoryImpl),
	}
}

type stateRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	countryRepository repository.CountryRepo

	pointsOfInterestRepositoryImpl repository.BaseRepo[*ent.PointsOfInterest, *do.PointsOfInterest]
	cityRepositoryImpl             repository.BaseRepo[*ent.City, *do.City]
	countryRepositoryImpl          repository.BaseRepo[*ent.Country, *do.Country]
}

func (s *stateRepositoryImpl) ToEntity(po *ent.State) *do.State {
	if po == nil {
		return nil
	}
	return &do.State{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		CountryID: po.CountryID,
		Poi:       s.pointsOfInterestRepositoryImpl.ToEntities(po.Edges.Poi),
		City:      s.cityRepositoryImpl.ToEntities(po.Edges.City),
		Country:   s.countryRepositoryImpl.ToEntity(po.Edges.Country),
	}
}

func (s *stateRepositoryImpl) ToEntities(pos []*ent.State) []*do.State {
	if len(pos) == 0 {
		return nil
	}

	list := make([]*do.State, 0, len(pos))
	for _, v := range pos {
		list = append(list, s.ToEntity(v))
	}
	return list
}

func (s *stateRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsState() {
		return s.countryRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := s.GetTx(ctx).State
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(state.Or(
			state.NameContains(keyword),
			state.NameCnContains(keyword),
			state.NameEnContains(keyword),
		)).
		WithCountry(func(query *ent.CountryQuery) {
			query.WithContinent()
		}).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := make([]*do.PointsOfInterest, 0, len(rows))
	for _, v := range rows {
		stateDo := s.ToEntity(v)
		pois = append(pois, &do.PointsOfInterest{
			State:     stateDo,
			Country:   stateDo.Country,
			Continent: stateDo.Country.Continent,
		})
	}
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := s.countryRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}

	return pois, nil
}

// ListStates 列出州/省，可选按国家过滤
func (r *stateRepositoryImpl) ListStates(ctx context.Context, params *bo.ListStatesParams) ([]*do.State, *bo.PaginationResult, error) {
	query := r.GetTx(ctx).State.Query()

	// 按国家过滤
	if params.CountryID > 0 {
		query = query.Where(state.CountryID(params.CountryID))
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

	// 加载关联的国家信息
	query = query.WithCountry()

	// 执行查询
	states, err := query.Order(ent.Asc(state.FieldName)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	// 判断是否有更多数据
	hasMore := false
	if len(states) > limit {
		hasMore = true
		states = states[:limit] // 去掉多查询的一条
	}

	// 计算下一页的偏移量
	nextOffset := params.Offset + len(states)

	// 转换为DO对象
	result := make([]*do.State, len(states))
	for i, s := range states {
		result[i] = r.ToEntity(s)
	}

	// 生成下一页令牌
	nextPageToken := bo.GenerateNextPageToken(nextOffset, hasMore)

	return result, &bo.PaginationResult{
		NextPageToken: nextPageToken,
		HasMore:       hasMore,
	}, nil
}

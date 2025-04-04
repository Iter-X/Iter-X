package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/iter-x/iter-x/internal/data/ent/city"
	"github.com/iter-x/iter-x/internal/data/ent/predicate"
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/pointsofinterest"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

type pointsOfInterestRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
	es     *data.Es

	cityRepository repository.CityRepo
}

func NewPointsOfInterest(d *data.Data, cityRepository repository.CityRepo, logger *zap.SugaredLogger) repository.PointsOfInterestRepo {
	return &pointsOfInterestRepositoryImpl{
		Tx:             d.Tx,
		es:             d.Es,
		logger:         logger.Named("repo.poi"),
		cityRepository: cityRepository,
	}
}

func (r *pointsOfInterestRepositoryImpl) ToEntity(po *ent.PointsOfInterest) *do.PointsOfInterest {
	if po == nil {
		return nil
	}
	return build.PointsOfInterestRepositoryImplToEntity(po)
}

func (r *pointsOfInterestRepositoryImpl) ToEntities(pos []*ent.PointsOfInterest) []*do.PointsOfInterest {
	if len(pos) == 0 {
		return nil
	}
	return build.PointsOfInterestRepositoryImplToEntities(pos)
}

const esIndex = "points_of_interest"

type (
	SearchResponse struct {
		Took     int         `json:"took"`
		TimedOut bool        `json:"timed_out"`
		Shards   ShardsInfo  `json:"_shards"`
		Hits     HitsWrapper `json:"hits"`
	}
	ShardsInfo struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	}
	HitsWrapper struct {
		Total    TotalHits `json:"total"`
		MaxScore float64   `json:"max_score"`
		Hits     []Hit     `json:"hits"`
	}
	TotalHits struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	}
	Hit struct {
		Index  string  `json:"_index"`
		ID     string  `json:"_id"`
		Score  float64 `json:"_score"`
		Source POI     `json:"_source"`
	}
	POI struct {
		ID                         string   `json:"id"`
		Name                       string   `json:"name"`
		NameEn                     string   `json:"name_en"`
		NameCn                     string   `json:"name_cn"`
		Description                string   `json:"description"`
		Address                    string   `json:"address"`
		Type                       string   `json:"type"`
		Category                   string   `json:"category"`
		Rating                     float64  `json:"rating"`
		RecommendedDurationMinutes int      `json:"recommended_duration_minutes"`
		CityID                     uint     `json:"city_id"`
		ContinentID                uint     `json:"continent_id"`
		CountryID                  uint     `json:"country_id"`
		StateID                    uint     `json:"state_id"`
		Location                   Location `json:"location"`
	}
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
)

func (r *pointsOfInterestRepositoryImpl) SearchPointsOfInterestByNamesFromES(ctx context.Context, names []string) ([]*do.PointsOfInterest, error) {
	if len(names) == 0 {
		r.logger.Warn("names cannot be empty")
		return nil, fmt.Errorf("names cannot be empty")
	}

	var shouldQueries []map[string]map[string]string
	for _, name := range names {
		shouldQueries = append(shouldQueries, map[string]map[string]string{
			"match": {"name": name},
		})
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should":               shouldQueries,
				"minimum_should_match": 1,
			},
		},
	}

	queryBody, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}

	res, err := r.es.Cli.Search(
		r.es.Cli.Search.WithContext(ctx),
		r.es.Cli.Search.WithIndex(esIndex),
		r.es.Cli.Search.WithBody(bytes.NewReader(queryBody)),
	)
	if err != nil {
		r.logger.Warnf("failed to search points of interest from elasticsearch: %v", err)
		return nil, nil
	}

	if res.StatusCode != 200 {
		r.logger.Warnf("failed to search points of interest from elasticsearch: %v", res.String())
		return nil, nil
	}

	var sr SearchResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&sr)
	if err != nil {
		r.logger.Warnf("failed to decode search response: %v", err)
		return nil, nil
	}

	pois := make([]*do.PointsOfInterest, 0, len(sr.Hits.Hits))
	for _, hit := range sr.Hits.Hits {
		id, err := uuid.Parse(hit.Source.ID)
		if err != nil {
			r.logger.Warnf("failed to parse poi id: %s", hit.Source.ID)
			continue
		}
		poi := &do.PointsOfInterest{
			ID:                         id,
			NameEn:                     hit.Source.NameEn,
			NameCn:                     hit.Source.NameCn,
			Description:                hit.Source.Description,
			Address:                    hit.Source.Address,
			Latitude:                   hit.Source.Location.Lat,
			Longitude:                  hit.Source.Location.Lon,
			Type:                       hit.Source.Type,
			Category:                   hit.Source.Category,
			Rating:                     float32(hit.Source.Rating),
			RecommendedDurationMinutes: int64(hit.Source.RecommendedDurationMinutes),
			CityID:                     hit.Source.CityID,
			StateID:                    hit.Source.StateID,
			CountryID:                  hit.Source.CountryID,
			ContinentID:                hit.Source.ContinentID,
			City:                       new(do.City),
			State:                      new(do.State),
			Country:                    new(do.Country),
			Continent:                  new(do.Continent),
		}
		pois = append(pois, poi)
	}

	return pois, nil
}

func (r *pointsOfInterestRepositoryImpl) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	if !params.IsPoi() {
		return r.cityRepository.SearchPointsOfInterest(ctx, params)
	}
	cli := r.GetTx(ctx).PointsOfInterest
	keyword := params.Keyword
	limit := params.Limit
	rows, err := cli.Query().
		Where(pointsofinterest.Or(
			pointsofinterest.NameLocalContains(keyword),
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
	pois := r.ToEntities(rows)
	otherRowLimit := limit - len(rows)
	if otherRowLimit > 0 && params.IsNext() {
		poiDos, err := r.cityRepository.SearchPointsOfInterest(ctx, params.DepthDec())
		if err != nil {
			return nil, err
		}
		pois = append(pois, poiDos...)
	}
	return pois, nil
}

func (r *pointsOfInterestRepositoryImpl) GetByCityNames(ctx context.Context, cityNames []string) ([]*do.PointsOfInterest, error) {
	if len(cityNames) == 0 {
		return nil, nil
	}
	cli := r.GetTx(ctx).City
	predicates := make([]predicate.City, 0, len(cityNames))
	for _, name := range cityNames {
		predicates = append(predicates, city.NameLocalContains(name))
	}
	cityRows, err := cli.Query().
		Where(city.Or(predicates...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	cityIds := make([]uint, 0, len(cityRows))
	for _, v := range cityRows {
		cityIds = append(cityIds, v.ID)
	}
	rows, err := r.GetTx(ctx).PointsOfInterest.Query().Select("id", "name").
		Where(pointsofinterest.CityIDIn(cityIds...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	pois := r.ToEntities(rows)
	return pois, nil
}

func (r *pointsOfInterestRepositoryImpl) GetTopPOIsByCity(ctx context.Context, cityIds []int32, limit int) ([]*do.PointsOfInterest, error) {
	if len(cityIds) == 0 {
		return nil, nil
	}

	// Convert []int32 to []uint
	uintCityIds := make([]uint, len(cityIds))
	for i, id := range cityIds {
		uintCityIds[i] = uint(id)
	}

	rows, err := r.GetTx(ctx).PointsOfInterest.Query().
		Where(pointsofinterest.CityIDIn(uintCityIds...)).
		Order(ent.Desc(pointsofinterest.FieldRating)).
		Limit(limit).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntities(rows), nil
}

func (r *pointsOfInterestRepositoryImpl) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*do.PointsOfInterest, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	rows, err := r.GetTx(ctx).PointsOfInterest.Query().
		Where(pointsofinterest.IDIn(ids...)).
		WithContinent().
		WithCountry().
		WithState().
		WithCity().
		All(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntities(rows), nil
}

func (r *pointsOfInterestRepositoryImpl) ListPOIs(ctx context.Context, params *bo.ListPOIsParams) ([]*do.PointsOfInterest, int64, error) {
	cli := r.GetTx(ctx).PointsOfInterest.Query()

	if params.CityId != nil {
		cli = cli.Where(pointsofinterest.CityID(uint(*params.CityId)))
	}

	if params.Keyword != nil && *params.Keyword != "" {
		cli = cli.Where(pointsofinterest.Or(
			pointsofinterest.NameLocalContains(*params.Keyword),
			pointsofinterest.NameCnContains(*params.Keyword),
			pointsofinterest.NameEnContains(*params.Keyword),
			pointsofinterest.DescriptionContains(*params.Keyword),
		))
	}

	total, err := cli.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if params.Pagination != nil {
		cli = cli.Offset(params.Pagination.GetOffset4Db()).Limit(params.Pagination.GetLimit4Db())
	}

	rows, err := cli.
		//WithCity().
		WithPoiFiles(func(query *ent.PointsOfInterestFilesQuery) {
			query.WithFile()
		}).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return r.ToEntities(rows), int64(total), nil
}

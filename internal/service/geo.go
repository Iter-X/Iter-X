package service

import (
	"context"

	geoV1 "github.com/iter-x/iter-x/internal/api/geo/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"go.uber.org/zap"
)

// GeoService 地理信息服务
type GeoService struct {
	geoV1.UnimplementedGeoServiceServer
	geo    *biz.Geo
	logger *zap.SugaredLogger
}

// NewGeoService 创建地理信息服务实例
func NewGeoService(geo *biz.Geo, logger *zap.SugaredLogger) *GeoService {
	return &GeoService{
		geo:    geo,
		logger: logger.Named("service.geo"),
	}
}

// ListContinents 列出所有大洲
func (s *GeoService) ListContinents(ctx context.Context, req *geoV1.ListContinentsRequest) (*geoV1.ListContinentsResponse, error) {
	return s.geo.ListContinents(ctx, req)
}

// ListCountries 列出国家，可选按大洲过滤
func (s *GeoService) ListCountries(ctx context.Context, req *geoV1.ListCountriesRequest) (*geoV1.ListCountriesResponse, error) {
	return s.geo.ListCountries(ctx, req)
}

// ListStates 列出州/省，可选按国家过滤
func (s *GeoService) ListStates(ctx context.Context, req *geoV1.ListStatesRequest) (*geoV1.ListStatesResponse, error) {
	return s.geo.ListStates(ctx, req)
}

// ListCities 列出城市，可选按州/省过滤
func (s *GeoService) ListCities(ctx context.Context, req *geoV1.ListCitiesRequest) (*geoV1.ListCitiesResponse, error) {
	return s.geo.ListCities(ctx, req)
}

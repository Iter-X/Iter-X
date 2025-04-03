package service

import (
	"context"

	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/service/build"
)

type PointsOfInterestService struct {
	poiV1.UnimplementedPointsOfInterestServiceServer
	pointsOfInterestBiz *biz.PointsOfInterest
}

func NewPointsOfInterest(pointsOfInterestBiz *biz.PointsOfInterest) *PointsOfInterestService {
	return &PointsOfInterestService{
		pointsOfInterestBiz: pointsOfInterestBiz,
	}
}

func (s *PointsOfInterestService) SearchPointsOfInterest(ctx context.Context, req *poiV1.SearchPointsOfInterestRequest) (*poiV1.SearchPointsOfInterestResponse, error) {
	params := &bo.SearchPointsOfInterestParams{
		Keyword:        req.GetKeyword(),
		Limit:          int(req.GetLimit()),
		GeographyLevel: req.GetGeographyLevel(),
	}
	pointsOfInterest, err := s.pointsOfInterestBiz.SearchPointsOfInterest(ctx, params.WithDepth(uint8(req.GetDepth())))
	if err != nil {
		return nil, err
	}
	return &poiV1.SearchPointsOfInterestResponse{
		PointsOfInterest: build.ToPointsOfInterestsProto(ctx, pointsOfInterest),
	}, nil
}

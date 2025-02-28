package service

import (
	"context"

	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/biz"
)

type PointsOfInterestService struct {
	poiV1.UnimplementedPointsOfInterestServiceServer
	pointsOfInterestBiz *biz.PointsOfInterest
}

func NewPointsOfInterestService(pointsOfInterestBiz *biz.PointsOfInterest) *PointsOfInterestService {
	return &PointsOfInterestService{
		pointsOfInterestBiz: pointsOfInterestBiz,
	}
}

func (s *PointsOfInterestService) SearchPointsOfInterest(ctx context.Context, req *poiV1.SearchPointsOfInterestRequest) (*poiV1.SearchPointsOfInterestResponse, error) {
	pointsOfInterest, err := s.pointsOfInterestBiz.SearchPointsOfInterest(ctx, req.Keyword, req.InitialCity)
	if err != nil {
		return nil, err
	}
	return &poiV1.SearchPointsOfInterestResponse{PointsOfInterest: pointsOfInterest}, nil
}

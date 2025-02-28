package service

import (
	"context"

	"github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/biz"
)

type PointsOfInterestService struct {
	v1.UnimplementedPointsOfInterestServiceServer
	biz *biz.PointsOfInterest
}

func NewPointsOfInterestService(biz *biz.PointsOfInterest) *PointsOfInterestService {
	return &PointsOfInterestService{
		biz: biz,
	}
}

func (s *PointsOfInterestService) SearchPointsOfInterest(ctx context.Context, req *v1.SearchPointsOfInterestRequest) (*v1.SearchPointsOfInterestResponse, error) {
	pointsOfInterest, err := s.biz.SearchPointsOfInterest(ctx, req.Keyword, req.InitialCity)
	if err != nil {
		return nil, err
	}
	return &v1.SearchPointsOfInterestResponse{PointsOfInterest: pointsOfInterest}, nil
}

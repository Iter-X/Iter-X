package biz

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/common/xerr"
)

type PointsOfInterest struct {
	pointsOfInterestRepo repository.PointsOfInterestRepo
}

func NewPointsOfInterest(pointsOfInterestRepo repository.PointsOfInterestRepo) *PointsOfInterest {
	return &PointsOfInterest{
		pointsOfInterestRepo: pointsOfInterestRepo,
	}
}

func (b *PointsOfInterest) SearchPointsOfInterest(ctx context.Context, params *bo.SearchPointsOfInterestParams) ([]*do.PointsOfInterest, error) {
	// TODO: Search points of interest by initial city at first, expand the search if no result found
	pois, err := b.pointsOfInterestRepo.SearchPointsOfInterest(ctx, params)
	if err != nil {
		return nil, xerr.ErrorSearchPoiFailed()
	}

	return pois, nil
}

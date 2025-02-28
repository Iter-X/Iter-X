package biz

import (
	"context"

	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
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

func (b *PointsOfInterest) SearchPointsOfInterest(ctx context.Context, keyword, initialCity string) ([]*poiV1.PointOfInterest, error) {
	// TODO: Search points of interest by initial city at first, expand the search if no result found
	pois, err := b.pointsOfInterestRepo.SearchPointsOfInterest(ctx, keyword, 5)
	if err != nil {
		return nil, xerr.ErrorSearchPoiFailed()
	}

	var res = make([]*poiV1.PointOfInterest, 0, len(pois))
	for _, poi := range pois {
		res = append(res, poi.ToPointsOfInterestProto())
	}
	return res, nil
}

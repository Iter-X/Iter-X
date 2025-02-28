package biz

import (
	"context"
	v1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/common/xerr"

	"github.com/iter-x/iter-x/internal/repo"
)

type PointsOfInterest struct {
	repo *repo.PointsOfInterest
}

func NewPointsOfInterest(repo *repo.PointsOfInterest) *PointsOfInterest {
	return &PointsOfInterest{
		repo: repo,
	}
}

func (b *PointsOfInterest) SearchPointsOfInterest(ctx context.Context, keyword, initialCity string) ([]*v1.PointOfInterest, error) {
	// TODO: Search points of interest by initial city at first, expand the search if no result found
	pois, err := b.repo.SearchPointsOfInterest(ctx, keyword, 5)
	if err != nil {
		return nil, xerr.ErrorSearchPoiFailed()
	}

	var res = make([]*v1.PointOfInterest, 0, len(pois))
	for _, poi := range pois {
		res = append(res, &v1.PointOfInterest{
			Id:                         poi.ID.String(),
			Name:                       poi.Name,
			NameEn:                     poi.NameEn,
			NameCn:                     poi.NameCn,
			Description:                poi.Description,
			Address:                    poi.Address,
			Latitude:                   poi.Latitude,
			Longitude:                  poi.Longitude,
			Type:                       poi.Type,
			Category:                   poi.Category,
			Rating:                     poi.Rating,
			RecommendedDurationMinutes: poi.RecommendedDurationMinutes,
			City:                       poi.Edges.City.Name,
			State:                      poi.Edges.State.Name,
			Country:                    poi.Edges.Country.Name,
		})
	}
	return res, nil
}

package build

import (
	"context"
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/pkg/util"
)

// ToPointsOfInterestProto copies the current object into a PointsOfInterest.
func ToPointsOfInterestProto(ctx context.Context, p *do.PointsOfInterest) *poiV1.PointOfInterest {
	if p == nil {
		return nil
	}
	return &poiV1.PointOfInterest{
		Id:                         p.ID.String(),
		Name:                       util.GetLocalizedName(ctx, p.NameEn, p.NameCn),
		NameEn:                     p.NameEn,
		NameCn:                     p.NameCn,
		Description:                p.Description,
		Address:                    p.Address,
		Latitude:                   p.Latitude,
		Longitude:                  p.Longitude,
		Type:                       p.Type,
		Category:                   p.Category,
		Rating:                     p.Rating,
		RecommendedDurationMinutes: p.RecommendedDurationMinutes,
		City:                       util.GetLocalizedName(ctx, p.City.NameEn, p.City.NameCn),
		State:                      util.GetLocalizedName(ctx, p.State.NameEn, p.State.NameCn),
		Country:                    util.GetLocalizedName(ctx, p.Country.NameEn, p.Country.NameCn),
	}
}

func ToPointsOfInterestsProto(ctx context.Context, ps []*do.PointsOfInterest) []*poiV1.PointOfInterest {
	if ps == nil {
		return nil
	}
	var pointsOfInterests []*poiV1.PointOfInterest
	for _, p := range ps {
		pointsOfInterests = append(pointsOfInterests, ToPointsOfInterestProto(ctx, p))
	}
	return pointsOfInterests
}

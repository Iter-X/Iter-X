package build

import (
	poiV1 "github.com/iter-x/iter-x/internal/api/poi/v1"
	"github.com/iter-x/iter-x/internal/biz/do"
)

// ToPointsOfInterestProto copies the current object into a PointsOfInterest.
func ToPointsOfInterestProto(p *do.PointsOfInterest) *poiV1.PointOfInterest {
	if p == nil {
		return nil
	}
	return &poiV1.PointOfInterest{
		Id:                         p.ID.String(),
		Name:                       p.Name,
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
		City:                       p.City.Name,
		State:                      p.State.Name,
		Country:                    p.Country.Name,
	}
}

func ToPointsOfInterestsProto(ps []*do.PointsOfInterest) []*poiV1.PointOfInterest {
	if ps == nil {
		return nil
	}
	var pointsOfInterests []*poiV1.PointOfInterest
	for _, p := range ps {
		pointsOfInterests = append(pointsOfInterests, ToPointsOfInterestProto(p))
	}
	return pointsOfInterests
}

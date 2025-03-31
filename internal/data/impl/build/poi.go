package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func PointsOfInterestRepositoryImplToEntity(po *ent.PointsOfInterest) *do.PointsOfInterest {
	if po == nil {
		return nil
	}
	return &do.PointsOfInterest{
		ID:                         po.ID,
		CreatedAt:                  po.CreatedAt,
		UpdatedAt:                  po.UpdatedAt,
		NameLocal:                  po.NameLocal,
		NameEn:                     po.NameEn,
		NameCn:                     po.NameCn,
		Description:                po.Description,
		Address:                    po.Address,
		Latitude:                   po.Latitude,
		Longitude:                  po.Longitude,
		Type:                       po.Type,
		Category:                   po.Category,
		Rating:                     po.Rating,
		RecommendedDurationMinutes: po.RecommendedDurationMinutes,
		CityID:                     po.CityID,
		StateID:                    po.StateID,
		CountryID:                  po.CountryID,
		ContinentID:                po.ContinentID,
		City:                       CityRepositoryImplToEntity(po.Edges.City),
		State:                      StateRepositoryImplToEntity(po.Edges.State),
		Country:                    CountryRepositoryImplToEntity(po.Edges.Country),
		Continent:                  ContinentRepositoryImplToEntity(po.Edges.Continent),
		DailyItinerary:             DailyItineraryRepositoryImplToEntities(po.Edges.DailyItinerary),
	}
}

func PointsOfInterestRepositoryImplToEntities(pois []*ent.PointsOfInterest) []*do.PointsOfInterest {
	if pois == nil {
		return nil
	}
	list := make([]*do.PointsOfInterest, 0, len(pois))
	for _, poi := range pois {
		list = append(list, PointsOfInterestRepositoryImplToEntity(poi))
	}
	return list
}

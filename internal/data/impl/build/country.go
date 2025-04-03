package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func CountryRepositoryImplToEntity(po *ent.Country) *do.Country {
	if po == nil {
		return nil
	}

	return &do.Country{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		NameLocal:   po.NameLocal,
		NameEn:      po.NameEn,
		NameCn:      po.NameCn,
		Code:        po.Code,
		ContinentID: po.ContinentID,
		Poi:         PointsOfInterestRepositoryImplToEntities(po.Edges.Poi),
		State:       StateRepositoryImplToEntities(po.Edges.State),
		Continent:   ContinentRepositoryImplToEntity(po.Edges.Continent),
		Image:       FileRepositoryImplToEntity(po.Edges.Image),
	}
}

func CountryRepositoryImplToEntities(pos []*ent.Country) []*do.Country {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.Country, 0, len(pos))
	for _, v := range pos {
		list = append(list, CountryRepositoryImplToEntity(v))
	}
	return list
}

package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func CityRepositoryImplToEntity(po *ent.City) *do.City {
	if po == nil {
		return nil
	}
	return &do.City{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		NameLocal: po.NameLocal,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		StateID:   po.StateID,
		Poi:       PointsOfInterestRepositoryImplToEntities(po.Edges.Poi),
		State:     StateRepositoryImplToEntity(po.Edges.State),
	}
}

func CityRepositoryImplToEntities(pos []*ent.City) []*do.City {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.City, 0, len(pos))
	for _, v := range pos {
		list = append(list, CityRepositoryImplToEntity(v))
	}
	return list
}

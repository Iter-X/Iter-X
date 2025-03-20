package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func StateRepositoryImplToEntity(po *ent.State) *do.State {
	if po == nil {
		return nil
	}
	return &do.State{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Name,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		CountryID: po.CountryID,
		Poi:       PointsOfInterestRepositoryImplToEntities(po.Edges.Poi),
		City:      CityRepositoryImplToEntities(po.Edges.City),
		Country:   CountryRepositoryImplToEntity(po.Edges.Country),
	}
}

func StateRepositoryImplToEntities(pos []*ent.State) []*do.State {
	if len(pos) == 0 {
		return nil
	}

	list := make([]*do.State, 0, len(pos))
	for _, v := range pos {
		list = append(list, StateRepositoryImplToEntity(v))
	}
	return list
}

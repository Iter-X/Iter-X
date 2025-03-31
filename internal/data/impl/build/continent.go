package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func ContinentRepositoryImplToEntity(po *ent.Continent) *do.Continent {
	if po == nil {
		return nil
	}
	return &do.Continent{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		NameLocal: po.NameLocal,
		NameEn:    po.NameEn,
		NameCn:    po.NameCn,
		Code:      po.Code,
		Poi:       PointsOfInterestRepositoryImplToEntities(po.Edges.Poi),
		Country:   CountryRepositoryImplToEntities(po.Edges.Country),
	}
}

func ContinentRepositoryImplToEntities(pos []*ent.Continent) []*do.Continent {
	if pos == nil {
		return nil
	}
	list := make([]*do.Continent, 0, len(pos))
	for _, v := range pos {
		list = append(list, ContinentRepositoryImplToEntity(v))
	}
	return list
}

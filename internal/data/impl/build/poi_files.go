package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func PoiFilesRepositoryImplToEntity(po *ent.PointsOfInterestFiles) *do.PointsOfInterestFiles {
	if po == nil {
		return nil
	}
	return &do.PointsOfInterestFiles{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		PoiID:     po.PoiID,
		FileID:    po.FileID,
		Poi:       PointsOfInterestRepositoryImplToEntity(po.Edges.Poi),
		File:      FileRepositoryImplToEntity(po.Edges.File),
	}
}

func PoiFilesRepositoryImplToEntities(pos []*ent.PointsOfInterestFiles) []*do.PointsOfInterestFiles {
	if pos == nil {
		return nil
	}
	list := make([]*do.PointsOfInterestFiles, 0, len(pos))
	for _, po := range pos {
		list = append(list, PoiFilesRepositoryImplToEntity(po))
	}
	return list
}

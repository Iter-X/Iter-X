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

func FileRepositoryImplToEntity(po *ent.File) *do.File {
	if po == nil {
		return nil
	}
	return &do.File{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		UserID:    po.UserID,
		Name:      po.Name,
		ObjectKey: po.ObjectKey,
		Size:      po.Size,
		URL:       po.URL,
		Star:      po.Star,
		Ext:       po.Ext,
		User:      AuthRepositoryImplToEntity(po.Edges.User),
		PoiFiles:  PoiFilesRepositoryImplToEntities(po.Edges.PoiFiles),
	}
}

func FileRepositoryImplToEntities(pos []*ent.File) []*do.File {
	if pos == nil {
		return nil
	}
	list := make([]*do.File, 0, len(pos))
	for _, po := range pos {
		list = append(list, FileRepositoryImplToEntity(po))
	}
	return list
}

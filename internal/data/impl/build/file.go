package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

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
		Ext:       po.Ext,
		User:      AuthRepositoryImplToEntity(po.Edges.User),
		PoiFiles:  PoiFilesRepositoryImplToEntities(po.Edges.PoiFiles),
	}
}

func FileRepositoryImplToEntities(pos []*ent.File) []*do.File {
	if len(pos) == 0 {
		return nil
	}
	list := make([]*do.File, 0, len(pos))
	for _, v := range pos {
		list = append(list, FileRepositoryImplToEntity(v))
	}
	return list
}

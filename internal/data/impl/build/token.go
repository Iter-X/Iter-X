package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func RefreshTokenImplToEntity(po *ent.RefreshToken) *do.RefreshToken {
	if po == nil {
		return nil
	}
	return &do.RefreshToken{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Token:     po.Token,
		ExpiresAt: po.ExpiresAt,
		UserID:    po.UserID,
		User:      AuthRepositoryImplToEntity(po.Edges.User),
	}
}

func RefreshTokenImplToEntities(pos []*ent.RefreshToken) []*do.RefreshToken {
	if pos == nil {
		return nil
	}
	list := make([]*do.RefreshToken, 0, len(pos))
	for _, v := range pos {
		list = append(list, RefreshTokenImplToEntity(v))
	}
	return list
}

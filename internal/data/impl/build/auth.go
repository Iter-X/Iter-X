package build

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/pkg/vobj"
)

func AuthRepositoryImplToEntity(po *ent.User) *do.User {
	if po == nil {
		return nil
	}
	return &do.User{
		ID:            po.ID,
		CreatedAt:     po.CreatedAt,
		UpdatedAt:     po.UpdatedAt,
		Status:        vobj.UserStatus(po.Status),
		Username:      po.Username,
		Password:      po.Password,
		Salt:          po.Salt,
		Nickname:      po.Nickname,
		Remark:        po.Remark,
		Phone:         po.Phone,
		Email:         po.Email,
		AvatarURL:     po.AvatarURL,
		RefreshTokens: RefreshTokenImplToEntities(po.Edges.RefreshToken),
		Trips:         TripRepositoryImplToEntities(po.Edges.Trip),
	}
}

func AuthRepositoryImplToEntities(pos []*ent.User) []*do.User {
	if pos == nil {
		return nil
	}
	list := make([]*do.User, 0, len(pos))
	for _, v := range pos {
		list = append(list, AuthRepositoryImplToEntity(v))
	}
	return list
}

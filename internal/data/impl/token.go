package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
)

func NewRefreshToken(d *data.Data, logger *zap.SugaredLogger) repository.TokenRepo {
	return &refreshTokenImpl{
		Tx:                 d.Tx,
		logger:             logger.Named("repo.refresh_token"),
		authRepositoryImpl: new(authRepositoryImpl),
	}
}

type refreshTokenImpl struct {
	*data.Tx
	logger *zap.SugaredLogger

	authRepositoryImpl repository.BaseRepo[*ent.User, *do.User]
}

func (r *refreshTokenImpl) ToEntity(po *ent.RefreshToken) *do.RefreshToken {
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
		User:      r.authRepositoryImpl.ToEntity(po.Edges.User),
	}
}

func (r *refreshTokenImpl) ToEntities(pos []*ent.RefreshToken) []*do.RefreshToken {
	if pos == nil {
		return nil
	}
	list := make([]*do.RefreshToken, 0, len(pos))
	for _, v := range pos {
		list = append(list, r.ToEntity(v))
	}
	return list
}

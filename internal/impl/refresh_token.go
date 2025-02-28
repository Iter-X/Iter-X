package impl

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"go.uber.org/zap"
)

func NewRefreshTokenRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.RefreshToken {
	return &refreshTokenRepositoryImpl{
		Tx:                 &Tx{cli: cli},
		logger:             logger.Named("repo.auth"),
		authRepositoryImpl: new(authRepositoryImpl),
	}
}

type refreshTokenRepositoryImpl struct {
	*Tx
	logger             *zap.SugaredLogger
	authRepositoryImpl *authRepositoryImpl
}

func (r *refreshTokenRepositoryImpl) ToEntity(po *ent.RefreshToken) *do.RefreshToken {
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

func (r *refreshTokenRepositoryImpl) ToEntities(pos []*ent.RefreshToken) []*do.RefreshToken {
	if pos == nil {
		return nil
	}
	list := make([]*do.RefreshToken, 0, len(pos))
	for _, v := range pos {
		list = append(list, r.ToEntity(v))
	}
	return list
}

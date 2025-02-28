package impl

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"go.uber.org/zap"
)

func NewRefreshTokenRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.RefreshToken {
	return &refreshTokenRepositoryImpl{
		Tx:     &Tx{cli: cli},
		logger: logger.Named("repo.auth"),
	}
}

type refreshTokenRepositoryImpl struct {
	*Tx
	logger *zap.SugaredLogger
}

func (r *refreshTokenRepositoryImpl) ToEntity(po *ent.RefreshToken) *do.RefreshToken {
	//TODO implement me
	panic("implement me")
}

func (r *refreshTokenRepositoryImpl) ToEntities(ts []*ent.RefreshToken) []*do.RefreshToken {
	//TODO implement me
	panic("implement me")
}

package impl

import (
	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewRefreshToken(d *data.Data, logger *zap.SugaredLogger) repository.TokenRepo {
	return &refreshTokenImpl{
		Tx:     d.Tx,
		logger: logger.Named("repo.refresh_token"),
	}
}

type refreshTokenImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (r *refreshTokenImpl) ToEntity(po *ent.RefreshToken) *do.RefreshToken {
	if po == nil {
		return nil
	}
	return build.RefreshTokenImplToEntity(po)
}

func (r *refreshTokenImpl) ToEntities(pos []*ent.RefreshToken) []*do.RefreshToken {
	if pos == nil {
		return nil
	}
	return build.RefreshTokenImplToEntities(pos)
}

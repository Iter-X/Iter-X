package impl

import (
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/impl/ent"
	"go.uber.org/zap"
)

func NewTripDailyRepository(cli *ent.Client, logger *zap.SugaredLogger) repository.DailyTrip[*ent.DailyTrip, *do.DailyTrip] {
	return &tripDailyRepositoryImpl{
		Tx:                 &Tx{cli: cli},
		logger:             logger.Named("repo.DailyTrip"),
		authRepositoryImpl: new(authRepositoryImpl),
	}
}

type tripDailyRepositoryImpl struct {
	*Tx
	logger             *zap.SugaredLogger
	authRepositoryImpl repository.Auth[*ent.User, *do.User]
}

func (t *tripDailyRepositoryImpl) ToEntity(po *ent.DailyTrip) *do.DailyTrip {
	//TODO implement me
	panic("implement me")
}

func (t *tripDailyRepositoryImpl) ToEntities(pos []*ent.DailyTrip) []*do.DailyTrip {
	//TODO implement me
	panic("implement me")
}

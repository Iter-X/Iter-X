package impl

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/pointsofinterestfiles"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

func NewPoiFiles(d *data.Data, logger *zap.SugaredLogger) repository.PoiFilesRepo {
	return &poiFilesRepositoryImpl{
		Tx:     d.Tx,
		logger: logger,
	}
}

type poiFilesRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (p *poiFilesRepositoryImpl) ToEntity(po *ent.PointsOfInterestFiles) *do.PointsOfInterestFiles {
	if po == nil {
		return nil
	}
	return build.PoiFilesRepositoryImplToEntity(po)
}

func (p *poiFilesRepositoryImpl) ToEntities(pos []*ent.PointsOfInterestFiles) []*do.PointsOfInterestFiles {
	return build.PoiFilesRepositoryImplToEntities(pos)
}

func (p *poiFilesRepositoryImpl) FindByPoiID(ctx context.Context, poiID uuid.UUID) ([]*do.PointsOfInterestFiles, error) {
	query := p.GetTx(ctx).PointsOfInterestFiles.Query().Where(pointsofinterestfiles.PoiID(poiID)).WithPoi().WithFile()
	row, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return build.PoiFilesRepositoryImplToEntities(row), nil
}

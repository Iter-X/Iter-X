package impl

import (
	"context"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/file"
	"github.com/iter-x/iter-x/internal/data/impl/build"
	"github.com/iter-x/iter-x/internal/helper/auth"
)

func NewFiles(d *data.Data, logger *zap.SugaredLogger) repository.FilesRepo {
	return &fileRepositoryImpl{
		Tx:     d.Tx,
		logger: logger,
	}
}

type fileRepositoryImpl struct {
	*data.Tx
	logger *zap.SugaredLogger
}

func (f *fileRepositoryImpl) ToEntity(po *ent.File) *do.File {
	if po == nil {
		return nil
	}
	return build.FileRepositoryImplToEntity(po)
}

func (f *fileRepositoryImpl) ToEntities(pos []*ent.File) []*do.File {
	return build.FileRepositoryImplToEntities(pos)
}

func (f *fileRepositoryImpl) Init(ctx context.Context, filename, objectKey string) error {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		return err
	}
	mutation := f.GetTx(ctx).File.Create().
		SetName(filename).
		SetObjectKey(objectKey).
		SetExt(filepath.Ext(filename)).
		SetUserID(claims.UID)
	return mutation.Exec(ctx)
}

func (f *fileRepositoryImpl) Complete(ctx context.Context, fileSize int64, complete *bo.CompleteUploadReply) error {
	if complete == nil {
		return nil
	}
	fileDo, err := f.FindByObjectKey(ctx, complete.Key)
	if err != nil {
		return err
	}
	complete.FileId = fileDo.ID
	mutation := f.GetTx(ctx).File.Update().
		Where(file.ID(fileDo.ID)).
		SetSize(uint(fileSize))
	return mutation.Exec(ctx)
}

func (f *fileRepositoryImpl) Delete(ctx context.Context, objectKey string) error {
	mutation := f.GetTx(ctx).File.Delete().Where(file.ObjectKey(objectKey))
	_, err := mutation.Exec(ctx)
	return err
}

func (f *fileRepositoryImpl) FindByID(ctx context.Context, id uint) (*do.File, error) {
	query := f.GetTx(ctx).File.Query().Where(file.ID(id))
	row, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return build.FileRepositoryImplToEntity(row), nil
}

func (f *fileRepositoryImpl) FindByObjectKey(ctx context.Context, objectKey string) (*do.File, error) {
	query := f.GetTx(ctx).File.Query().Where(file.ObjectKey(objectKey))
	row, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return build.FileRepositoryImplToEntity(row), nil
}

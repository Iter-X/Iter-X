package repository

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/data/ent"
)

type Files[T *ent.File, R *do.File] interface {
	BaseRepo[T, R]

	Init(ctx context.Context, filename, objectKey string) error
	Complete(ctx context.Context, fileSize int64, complete *bo.CompleteUploadReply) error
	Delete(ctx context.Context, objectKey string) error
	FindByID(ctx context.Context, id uint) (R, error)
	FindByObjectKey(ctx context.Context, objectKey string) (R, error)
}

type FilesRepo = Files[*ent.File, *do.File]

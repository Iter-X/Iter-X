package repository

import (
	"context"
	"io"

	"github.com/iter-x/iter-x/internal/biz/bo"
)

type StorageRepo interface {
	Upload(ctx context.Context, objectKey string, reader io.Reader) error
	Delete(ctx context.Context, objectKey string) error
	CompleteUpload(ctx context.Context, objectKey string) (*bo.CompleteUploadReply, error)
}

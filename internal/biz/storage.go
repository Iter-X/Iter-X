package biz

import (
	"context"
	"time"

	"github.com/iter-x/iter-x/internal/biz/bo"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/helper/auth"
	"github.com/iter-x/iter-x/pkg/storage"
)

func NewStorage(d *data.Data) *Storage {
	return &Storage{
		storage: d.Storage,
	}
}

type Storage struct {
	storage storage.FileManager
}

func (s *Storage) InitUpload(ctx context.Context, request *bo.InitUploadRequest) (*bo.InitUploadReply, error) {
	claims, err := auth.ExtractClaims(ctx)
	if err != nil {
		//claims = new(auth.Claims)
		//claims.UID = uuid.New()
		return nil, err
	}
	uploadResult, err := s.storage.InitiateMultipartUpload(request.Filename, claims.UID.String())
	if err != nil {
		return nil, err
	}
	return &bo.InitUploadReply{
		UploadID:   uploadResult.UploadID,
		BucketName: uploadResult.BucketName,
		ObjectKey:  uploadResult.ObjectKey,
	}, nil
}

func (s *Storage) GenerateUploadPartURL(_ context.Context, request *bo.GenerateUploadPartURLRequest) (*bo.GenerateUploadPartURLReply, error) {
	// TODO expiration time should be configurable
	uploadPartURL, err := s.storage.GenerateUploadPartURL(request.UploadID, request.ObjectKey, request.PartNumber, time.Minute*15)
	if err != nil {
		return nil, err
	}
	return &bo.GenerateUploadPartURLReply{
		UploadID:       uploadPartURL.UploadID,
		BucketName:     uploadPartURL.BucketName,
		ObjectKey:      uploadPartURL.ObjectKey,
		PartNumber:     uploadPartURL.PartNumber,
		UploadURL:      uploadPartURL.UploadURL,
		ExpirationTime: uploadPartURL.ExpirationTime,
	}, nil
}

func (s *Storage) CompleteUpload(_ context.Context, request *bo.CompleteUploadRequest) (*bo.CompleteUploadReply, error) {
	parts := make([]storage.UploadPart, 0, len(request.Parts))
	for _, part := range request.Parts {
		parts = append(parts, storage.UploadPart{
			ETag:       part.ETag,
			PartNumber: part.PartNumber,
		})
	}
	completeMultipartUpload, err := s.storage.CompleteMultipartUpload(request.UploadID, request.ObjectKey, parts)
	if err != nil {
		return nil, err
	}
	return &bo.CompleteUploadReply{
		Bucket:     completeMultipartUpload.Bucket,
		ETag:       completeMultipartUpload.ETag,
		Expiration: completeMultipartUpload.Expiration,
		Key:        completeMultipartUpload.Key,
		Location:   completeMultipartUpload.Location,
		PrivateURL: completeMultipartUpload.PrivateURL,
		PublicURL:  completeMultipartUpload.PublicURL,
	}, nil
}

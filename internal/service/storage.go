package service

import (
	"context"

	storageV1 "github.com/iter-x/iter-x/internal/api/storage/v1"
	"github.com/iter-x/iter-x/internal/biz"
	"github.com/iter-x/iter-x/internal/biz/bo"
)

var _ storageV1.StorageServer = (*Storage)(nil)

func NewStorage(storageBiz *biz.Storage) *Storage {
	return &Storage{
		storageBiz: storageBiz,
	}
}

type Storage struct {
	storageV1.UnimplementedStorageServer

	storageBiz *biz.Storage
}

func (s *Storage) InitUpload(ctx context.Context, request *storageV1.InitUploadRequest) (*storageV1.InitUploadReply, error) {
	params := &bo.InitUploadRequest{Filename: request.GetFilename()}
	reply, err := s.storageBiz.InitUpload(ctx, params)
	if err != nil {
		return nil, err
	}
	return &storageV1.InitUploadReply{
		UploadId:   reply.UploadID,
		ObjectKey:  reply.ObjectKey,
		BucketName: reply.BucketName,
	}, nil
}

func (s *Storage) GenerateUploadPartURL(ctx context.Context, request *storageV1.GenerateUploadPartURLRequest) (*storageV1.GenerateUploadPartURLReply, error) {
	params := &bo.GenerateUploadPartURLRequest{
		UploadID:   request.GetUploadId(),
		ObjectKey:  request.GetObjectKey(),
		PartNumber: int(request.GetPartNumber()),
	}
	reply, err := s.storageBiz.GenerateUploadPartURL(ctx, params)
	if err != nil {
		return nil, err
	}
	return &storageV1.GenerateUploadPartURLReply{
		UploadId:       reply.UploadID,
		ObjectKey:      reply.ObjectKey,
		BucketName:     reply.BucketName,
		UploadUrl:      reply.UploadURL,
		PartNumber:     int32(reply.PartNumber),
		ExpirationTime: reply.ExpirationTime,
	}, nil
}

func (s *Storage) CompleteUpload(ctx context.Context, request *storageV1.CompleteUploadRequest) (*storageV1.CompleteUploadReply, error) {
	params := &bo.CompleteUploadRequest{
		UploadID:  request.GetUploadId(),
		ObjectKey: request.GetObjectKey(),
		Parts:     make([]bo.UploadPart, 0),
		FileSize:  request.GetFileSize(),
	}
	for _, part := range request.GetParts() {
		params.Parts = append(params.Parts, bo.UploadPart{
			ETag:       part.GetEtag(),
			PartNumber: int(part.GetPartNumber()),
		})
	}
	reply, err := s.storageBiz.CompleteUpload(ctx, params)
	if err != nil {
		return nil, err
	}
	return &storageV1.CompleteUploadReply{
		Location:   reply.Location,
		Bucket:     reply.Bucket,
		Key:        reply.Key,
		Etag:       reply.ETag,
		PrivateUrl: reply.PrivateURL,
		PublicUrl:  reply.PublicURL,
		Expiration: reply.Expiration,
		FileId:     uint64(reply.FileId),
	}, nil
}

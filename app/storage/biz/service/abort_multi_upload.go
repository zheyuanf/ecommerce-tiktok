package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type AbortMultiUploadService struct {
	ctx context.Context
} // NewAbortMultiUploadService new AbortMultiUploadService
func NewAbortMultiUploadService(ctx context.Context) *AbortMultiUploadService {
	return &AbortMultiUploadService{ctx: ctx}
}

// Run create note info
func (s *AbortMultiUploadService) Run(req *storage.AbortMultiUploadRequest) (resp *storage.AbortMultiUploadResponse, err error) {
	// abort multi upload
	err = store.MinIOCore.AbortMultipartUpload(s.ctx, store.BucketName, req.FileName, req.UploadId)
	return
}

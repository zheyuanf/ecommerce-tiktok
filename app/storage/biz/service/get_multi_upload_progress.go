package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type GetMultiUploadProgressService struct {
	ctx context.Context
} // NewGetMultiUploadProgressService new GetMultiUploadProgressService
func NewGetMultiUploadProgressService(ctx context.Context) *GetMultiUploadProgressService {
	return &GetMultiUploadProgressService{ctx: ctx}
}

// Run create note info
func (s *GetMultiUploadProgressService) Run(req *storage.GetMultiUploadProgressRequest) (resp *storage.GetMultiUploadProgressResponse, err error) {
	// 获取已上传的分片信息
	res, err := store.MinIOCore.ListObjectParts(s.ctx, store.BucketName, req.FileName, req.UploadId, 0, 1024*1024)
	if err != nil {
		klog.Errorf("ListMultipartUploads error: %v", err)
		return
	}
	chunks := []int32{}
	for _, ul := range res.ObjectParts {
		chunks = append(chunks, int32(ul.PartNumber))
	}
	resp = &storage.GetMultiUploadProgressResponse{
		UploadedChunks: chunks,
	}
	return
}

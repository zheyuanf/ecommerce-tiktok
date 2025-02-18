package service

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type MergeFileChunksService struct {
	ctx context.Context
} // NewMergeFileChunksService new MergeFileChunksService
func NewMergeFileChunksService(ctx context.Context) *MergeFileChunksService {
	return &MergeFileChunksService{ctx: ctx}
}

// Run create note info
func (s *MergeFileChunksService) Run(req *storage.MergeFileChunksRequest) (resp *storage.MergeFileChunksResponse, err error) {
	// 获取已上传的分片信息
	res, err := store.MinIOCore.ListObjectParts(s.ctx, store.BucketName, req.FileName, req.UploadId, 0, 1024*1024)
	if err != nil {
		klog.Errorf("ListMultipartUploads error: %v", err)
		return
	}
	// 检查分片数量是否正确
	if len(res.ObjectParts) != int(req.TotalChunks) {
		klog.Errorf("file chunks not match: %v != %v", len(res.ObjectParts), req.TotalChunks)
		err = errors.New("file chunks not match")
		return
	}

	// 合并分片
	parts := make([]minio.CompletePart, len(res.ObjectParts))
	for i, op := range res.ObjectParts {
		parts[i] = minio.CompletePart{
			PartNumber: op.PartNumber,
			ETag:       op.ETag,
		}
	}
	uploadInfo, err := store.MinIOCore.CompleteMultipartUpload(
		s.ctx,
		store.BucketName,
		req.FileName,
		req.UploadId,
		parts,
		minio.PutObjectOptions{ContentType: req.FileType},
	)
	if err != nil {
		klog.Errorf("complete multipart upload: %v", err)
		return
	}
	klog.Infof("merge file parts success: %v", uploadInfo)
	return
}

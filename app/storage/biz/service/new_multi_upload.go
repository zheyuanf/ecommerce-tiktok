package service

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type NewMultiUploadService struct {
	ctx context.Context
} // NewNewMultiUploadService new NewMultiUploadService
func NewNewMultiUploadService(ctx context.Context) *NewMultiUploadService {
	return &NewMultiUploadService{ctx: ctx}
}

// Run create note info
func (s *NewMultiUploadService) Run(req *storage.NewMultiUploadRequest) (resp *storage.NewMultiUploadResponse, err error) {
	// 新建一个 multipart upload
	uploadId, err := store.MinIOCore.NewMultipartUpload(s.ctx, store.BucketName, req.FileName, minio.PutObjectOptions{ContentType: req.FileType})
	if err != nil {
		klog.Errorf("NewMultipartUpload error: %v", err)
		return
	}
	// 生成每个分片的 presigned URL
	presignedUrls := []*storage.URL{}
	for i := 1; i <= int(req.TotalChunks); i++ {
		// 设置查询参数
		params := url.Values{}
		params.Set("uploadId", uploadId)
		params.Set("partNumber", strconv.Itoa(i))
		url, err := store.MinIOClient.PresignHeader(
			s.ctx,
			"PUT",
			store.BucketName,
			req.FileName,
			time.Hour*24,
			params, http.Header{},
		)
		if err != nil {
			klog.Errorf("PresignedPutObject error: %v", err)
			return nil, err
		}
		presignedUrls = append(presignedUrls, &storage.URL{
			Url:         url.String(),
			ChunkNumber: int32(i),
		})
	}
	resp = &storage.NewMultiUploadResponse{
		UploadId:      uploadId,
		PresignedUrls: presignedUrls,
	}
	return
}

package service

import (
	"bytes"
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type UploadFileService struct {
	ctx context.Context
} // NewUploadFileService new UploadFileService
func NewUploadFileService(ctx context.Context) *UploadFileService {
	return &UploadFileService{ctx: ctx}
}

// Run create note info
func (s *UploadFileService) Run(req *storage.UploadFileRequest) (resp *storage.UploadFileResponse, err error) {
	// TODO: 1. 校验和
	checkSum := req.Checksum
	if checkSum != store.GetFileChecksum(req.FileData) {
		klog.Errorf("checksum error")
		err = errors.New("checksum error")
		return
	}
	// 2. 上传文件
	uploadInfo, err := store.MinIOClient.PutObject(s.ctx, store.BucketName, req.FileName, bytes.NewReader(req.FileData),
		int64(len(req.FileData)), minio.PutObjectOptions{ContentType: req.FileType})
	if err != nil {
		klog.Errorf("put object error: %v", err)
		return
	}
	klog.Infof("upload file success: %v", uploadInfo)
	return
}

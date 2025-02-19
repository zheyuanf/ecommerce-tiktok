package main

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/biz/service"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

// FileStorageServiceImpl implements the last service interface defined in the IDL.
type FileStorageServiceImpl struct{}

// MergeFileChunks implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) MergeFileChunks(ctx context.Context, req *storage.MergeFileChunksRequest) (resp *storage.MergeFileChunksResponse, err error) {
	resp, err = service.NewMergeFileChunksService(ctx).Run(req)

	return resp, err
}

// UploadFile implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) UploadFile(ctx context.Context, req *storage.UploadFileRequest) (resp *storage.UploadFileResponse, err error) {
	resp, err = service.NewUploadFileService(ctx).Run(req)

	return resp, err
}

// NewMultiUpload implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) NewMultiUpload(ctx context.Context, req *storage.NewMultiUploadRequest) (resp *storage.NewMultiUploadResponse, err error) {
	resp, err = service.NewNewMultiUploadService(ctx).Run(req)

	return resp, err
}

// AbortMultiUpload implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) AbortMultiUpload(ctx context.Context, req *storage.AbortMultiUploadRequest) (resp *storage.AbortMultiUploadResponse, err error) {
	resp, err = service.NewAbortMultiUploadService(ctx).Run(req)

	return resp, err
}

// GetMultiUploadProgress implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) GetMultiUploadProgress(ctx context.Context, req *storage.GetMultiUploadProgressRequest) (resp *storage.GetMultiUploadProgressResponse, err error) {
	resp, err = service.NewGetMultiUploadProgressService(ctx).Run(req)

	return resp, err
}

// GetDownloadUrl implements the FileStorageServiceImpl interface.
func (s *FileStorageServiceImpl) GetDownloadUrl(ctx context.Context, req *storage.GetDownloadUrlRequest) (resp *storage.GetDownloadUrlResponse, err error) {
	resp, err = service.NewGetDownloadUrlService(ctx).Run(req)

	return resp, err
}

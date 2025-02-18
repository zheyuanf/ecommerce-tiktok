package storage

import (
	"context"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func UploadFile(ctx context.Context, req *storage.UploadFileRequest, callOptions ...callopt.Option) (resp *storage.UploadFileResponse, err error) {
	resp, err = defaultClient.UploadFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UploadFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func NewMultiUpload(ctx context.Context, req *storage.NewMultiUploadRequest, callOptions ...callopt.Option) (resp *storage.NewMultiUploadResponse, err error) {
	resp, err = defaultClient.NewMultiUpload(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "NewMultiUpload call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AbortMultiUpload(ctx context.Context, req *storage.AbortMultiUploadRequest, callOptions ...callopt.Option) (resp *storage.AbortMultiUploadResponse, err error) {
	resp, err = defaultClient.AbortMultiUpload(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AbortMultiUpload call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetMultiUploadProgress(ctx context.Context, req *storage.GetMultiUploadProgressRequest, callOptions ...callopt.Option) (resp *storage.GetMultiUploadProgressResponse, err error) {
	resp, err = defaultClient.GetMultiUploadProgress(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetMultiUploadProgress call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MergeFileChunks(ctx context.Context, req *storage.MergeFileChunksRequest, callOptions ...callopt.Option) (resp *storage.MergeFileChunksResponse, err error) {
	resp, err = defaultClient.MergeFileChunks(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MergeFileChunks call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetDownloadUrl(ctx context.Context, req *storage.GetDownloadUrlRequest, callOptions ...callopt.Option) (resp *storage.GetDownloadUrlResponse, err error) {
	resp, err = defaultClient.GetDownloadUrl(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetDownloadUrl call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

package storage

import (
	"context"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"

	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage/filestorageservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() filestorageservice.Client
	Service() string
	UploadFile(ctx context.Context, Req *storage.UploadFileRequest, callOptions ...callopt.Option) (r *storage.UploadFileResponse, err error)
	NewMultiUpload(ctx context.Context, Req *storage.NewMultiUploadRequest, callOptions ...callopt.Option) (r *storage.NewMultiUploadResponse, err error)
	AbortMultiUpload(ctx context.Context, Req *storage.AbortMultiUploadRequest, callOptions ...callopt.Option) (r *storage.AbortMultiUploadResponse, err error)
	GetMultiUploadProgress(ctx context.Context, Req *storage.GetMultiUploadProgressRequest, callOptions ...callopt.Option) (r *storage.GetMultiUploadProgressResponse, err error)
	MergeFileChunks(ctx context.Context, Req *storage.MergeFileChunksRequest, callOptions ...callopt.Option) (r *storage.MergeFileChunksResponse, err error)
	GetDownloadUrl(ctx context.Context, Req *storage.GetDownloadUrlRequest, callOptions ...callopt.Option) (r *storage.GetDownloadUrlResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := filestorageservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient filestorageservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() filestorageservice.Client {
	return c.kitexClient
}

func (c *clientImpl) UploadFile(ctx context.Context, Req *storage.UploadFileRequest, callOptions ...callopt.Option) (r *storage.UploadFileResponse, err error) {
	return c.kitexClient.UploadFile(ctx, Req, callOptions...)
}

func (c *clientImpl) NewMultiUpload(ctx context.Context, Req *storage.NewMultiUploadRequest, callOptions ...callopt.Option) (r *storage.NewMultiUploadResponse, err error) {
	return c.kitexClient.NewMultiUpload(ctx, Req, callOptions...)
}

func (c *clientImpl) AbortMultiUpload(ctx context.Context, Req *storage.AbortMultiUploadRequest, callOptions ...callopt.Option) (r *storage.AbortMultiUploadResponse, err error) {
	return c.kitexClient.AbortMultiUpload(ctx, Req, callOptions...)
}

func (c *clientImpl) GetMultiUploadProgress(ctx context.Context, Req *storage.GetMultiUploadProgressRequest, callOptions ...callopt.Option) (r *storage.GetMultiUploadProgressResponse, err error) {
	return c.kitexClient.GetMultiUploadProgress(ctx, Req, callOptions...)
}

func (c *clientImpl) MergeFileChunks(ctx context.Context, Req *storage.MergeFileChunksRequest, callOptions ...callopt.Option) (r *storage.MergeFileChunksResponse, err error) {
	return c.kitexClient.MergeFileChunks(ctx, Req, callOptions...)
}

func (c *clientImpl) GetDownloadUrl(ctx context.Context, Req *storage.GetDownloadUrlRequest, callOptions ...callopt.Option) (r *storage.GetDownloadUrlResponse, err error) {
	return c.kitexClient.GetDownloadUrl(ctx, Req, callOptions...)
}

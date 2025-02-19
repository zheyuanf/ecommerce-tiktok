package service

import (
	"context"
	"net/url"
	"time"

	"github.com/zheyuanf/ecommerce-tiktok/app/storage/infra/store"
	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type GetDownloadUrlService struct {
	ctx context.Context
} // NewGetDownloadUrlService new GetDownloadUrlService
func NewGetDownloadUrlService(ctx context.Context) *GetDownloadUrlService {
	return &GetDownloadUrlService{ctx: ctx}
}

// Run create note info
func (s *GetDownloadUrlService) Run(req *storage.GetDownloadUrlRequest) (resp *storage.GetDownloadUrlResponse, err error) {
	// 获取文件下载链接，过期时间设置为24小时
	u, err := store.MinIOClient.PresignedGetObject(s.ctx, store.BucketName, req.FileName, time.Hour*24, url.Values{})
	if err != nil {
		return
	}
	resp = &storage.GetDownloadUrlResponse{
		DownloadUrl: u.String(),
	}
	return
}

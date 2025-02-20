package service

import (
	"context"

	storage "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/storage"
)

type GetFileUploadStatusService struct {
	ctx context.Context
} // NewGetFileUploadStatusService new GetFileUploadStatusService
func NewGetFileUploadStatusService(ctx context.Context) *GetFileUploadStatusService {
	return &GetFileUploadStatusService{ctx: ctx}
}

// Run create note info
func (s *GetFileUploadStatusService) Run(req *storage.GetFileUploadStatusRequest) (resp *storage.GetFileUploadStatusResponse, err error) {
	// Finish your business logic.

	return
}

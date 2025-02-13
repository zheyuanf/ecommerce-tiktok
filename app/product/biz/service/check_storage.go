package service

import (
	"context"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type CheckStorageService struct {
	ctx context.Context
} // NewCheckStorageService new CheckStorageService
func NewCheckStorageService(ctx context.Context) *CheckStorageService {
	return &CheckStorageService{ctx: ctx}
}

// Run create note info
func (s *CheckStorageService) Run(req *product.CheckStorageReq) (resp *product.CheckStorageResp, err error) {
	// Finish your business logic.

	return
}

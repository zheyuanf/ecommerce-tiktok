package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

type GetStockService struct {
	ctx context.Context
} // NewGetStockService new GetStockService
func NewGetStockService(ctx context.Context) *GetStockService {
	return &GetStockService{ctx: ctx}
}

// Run create note info
func (s *GetStockService) Run(req *inventory.GetStockRequest) (resp *inventory.GetStockResponse, err error) {
	ivt, err := model.QueryInventory(s.ctx, mysql.DB, int(req.ProductId))
	if err != nil {
		return
	}
	resp = &inventory.GetStockResponse{
		ProductId: int32(ivt.ProductID),
		Stock:     ivt.Stock - ivt.LockedStock,
	}
	return
}

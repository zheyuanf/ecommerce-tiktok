package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

type DecreaseStockService struct {
	ctx context.Context
} // NewDecreaseStockService new DecreaseStockService
func NewDecreaseStockService(ctx context.Context) *DecreaseStockService {
	return &DecreaseStockService{ctx: ctx}
}

// Run create note info
func (s *DecreaseStockService) Run(req *inventory.DecreaseStockRequest) (resp *inventory.DecreaseStockResponse, err error) {
	err = model.DecreaseStock(s.ctx, mysql.DB, req.ProductId, req.Quantity)
	if err != nil {
		return
	}
	resp = &inventory.DecreaseStockResponse{
		Success: true,
	}
	return
}

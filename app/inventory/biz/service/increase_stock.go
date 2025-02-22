package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

type IncreaseStockService struct {
	ctx context.Context
} // NewIncreaseStockService new IncreaseStockService
func NewIncreaseStockService(ctx context.Context) *IncreaseStockService {
	return &IncreaseStockService{ctx: ctx}
}

// Run create note info
func (s *IncreaseStockService) Run(req *inventory.IncreaseStockRequest) (resp *inventory.IncreaseStockResponse, err error) {
	err = model.IncreaseStock(s.ctx, mysql.DB, req.ProductId, req.Quantity)
	if err != nil {
		return
	}
	resp = &inventory.IncreaseStockResponse{
		Success: true,
	}
	return
}

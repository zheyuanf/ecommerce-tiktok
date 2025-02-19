package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

type BatchGetStockService struct {
	ctx context.Context
} // NewBatchGetStockService new BatchGetStockService
func NewBatchGetStockService(ctx context.Context) *BatchGetStockService {
	return &BatchGetStockService{ctx: ctx}
}

// Run create note info
func (s *BatchGetStockService) Run(req *inventory.BatchGetStockRequest) (resp *inventory.BatchGetStockResponse, err error) {
	ivts, err := model.BatchQueryInventory(s.ctx, mysql.DB, req.ProductIds)
	if err != nil {
		return
	}
	stockInfos := make([]*inventory.StockInfo, len(ivts))
	for i, ivt := range ivts {
		stockInfos[i] = &inventory.StockInfo{
			ProductId: ivt.ProductID,
			Stock:     ivt.Stock - ivt.LockedStock,
		}
	}
	return
}

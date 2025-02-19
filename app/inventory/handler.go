package main

import (
	"context"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/service"
)

// InventoryServiceImpl implements the last service interface defined in the IDL.
type InventoryServiceImpl struct{}

// GetStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) GetStock(ctx context.Context, req *inventory.GetStockRequest) (resp *inventory.GetStockResponse, err error) {
	resp, err = service.NewGetStockService(ctx).Run(req)

	return resp, err
}

// DecreaseStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) DecreaseStock(ctx context.Context, req *inventory.DecreaseStockRequest) (resp *inventory.DecreaseStockResponse, err error) {
	resp, err = service.NewDecreaseStockService(ctx).Run(req)

	return resp, err
}

// IncreaseStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) IncreaseStock(ctx context.Context, req *inventory.IncreaseStockRequest) (resp *inventory.IncreaseStockResponse, err error) {
	resp, err = service.NewIncreaseStockService(ctx).Run(req)

	return resp, err
}

// BatchGetStock implements the InventoryServiceImpl interface.
func (s *InventoryServiceImpl) BatchGetStock(ctx context.Context, req *inventory.BatchGetStockRequest) (resp *inventory.BatchGetStockResponse, err error) {
	resp, err = service.NewBatchGetStockService(ctx).Run(req)

	return resp, err
}

package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

func MockData(ctx context.Context) {
	err := model.CreateInventory(ctx, mysql.DB, &model.Inventory{
		ProductID:   1,
		Stock:       100,
		LockedStock: 30,
	})
	if err != nil {
		panic(err)
	}
}

func TestGetStock_Run(t *testing.T) {
	mysql.MockInit()
	ctx := context.Background()
	MockData(ctx)
	s := NewGetStockService(ctx)
	// init req and assert value

	req := &inventory.GetStockRequest{
		ProductId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}

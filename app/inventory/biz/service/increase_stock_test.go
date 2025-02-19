package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

func TestIncreaseStock_Run(t *testing.T) {
	ctx := context.Background()
	mysql.MockInit()
	s := NewIncreaseStockService(ctx)

	req := &inventory.IncreaseStockRequest{
		ProductId: 1,
		Quantity:  20,
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}

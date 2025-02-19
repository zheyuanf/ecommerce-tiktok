package service

import (
	"context"
	"testing"

	inventory "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/inventory"
)

func TestBatchGetStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewBatchGetStockService(ctx)
	// init req and assert value

	req := &inventory.BatchGetStockRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

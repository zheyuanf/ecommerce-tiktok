package service

import (
	"context"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
	"testing"
)

func TestCheckStorage_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckStorageService(ctx)
	// init req and assert value

	req := &product.CheckStorageReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

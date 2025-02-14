package service

import (
	"context"
	"testing"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

func TestUpdateProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateProductService(ctx)
	// init req and assert value

	req := &product.UpdateProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

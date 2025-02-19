package service

import (
	"context"
	"testing"

	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

func TestCreateCategory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateCategoryService(ctx)
	// init req and assert value

	req := &product.CreateCategoryReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

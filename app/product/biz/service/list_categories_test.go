package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

func TestListCategories_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListCategoriesService(ctx)
	mysql.MockInit()
	MockDataUpdate(ctx)
	// init req and assert value

	req := &product.ListCategoriesReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

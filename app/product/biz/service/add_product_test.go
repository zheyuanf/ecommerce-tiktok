package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

func TestAddProduct_Run(t *testing.T) {
	ctx := context.Background()
	mysql.MockInit()
	s := NewAddProductService(ctx)
	// init req and assert value

	q := model.NewCategoryQuery(ctx, mysql.DB)
	q.CreateCategory(&model.Category{
		Name:        "cname1",
		Description: "test description",
	})

	req := &product.AddProductReq{
		Name:        "test",
		Price:       100,
		Description: "test description",
		Categories:  []string{"cname1"},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

func MockData(ctx context.Context) {
	category := &model.Category{
		Name:        "category1",
		Description: "category1 description",
	}
	model.NewCategoryQuery(ctx, mysql.DB).CreateCategory(category)
	model.NewAddProduct(ctx, mysql.DB).AddProduct(&model.Product{
		Name:        "product1",
		Description: "product1 description",
		Picture:     "product1 picture",
		Categories:  []model.Category{*category},
		Price:       100,
	})
}

func TestDeleteProduct_Run(t *testing.T) {
	mysql.MockInit()
	ctx := context.Background()
	MockData(ctx)
	s := NewDeleteProductService(ctx)
	// init req and assert value

	req := &product.DeleteProductReq{
		Id: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	c, err := model.GetProductsByCategoryName(mysql.DB, s.ctx, req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = &product.ListProductsResp{}
	for _, v1 := range c {
		for _, v := range v1.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(v.ID),
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
				Storage:     v.Storage,
			})
		}
	}
	return resp, nil
}

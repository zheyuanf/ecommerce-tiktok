package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type ListCategoriesService struct {
	ctx context.Context
} // NewListCategoriesService new ListCategoriesService
func NewListCategoriesService(ctx context.Context) *ListCategoriesService {
	return &ListCategoriesService{ctx: ctx}
}

// Run create note info
func (s *ListCategoriesService) Run(req *product.ListCategoriesReq) (resp *product.ListCategoriesResp, err error) {
	categories, err := model.NewCategoryQuery(s.ctx, mysql.DB).ListCategories()
	if err != nil {
		return
	}
	names := make([]string, len(categories))
	for i, category := range categories {
		names[i] = category.Name
	}
	return &product.ListCategoriesResp{
		Categories: names,
	}, nil
}

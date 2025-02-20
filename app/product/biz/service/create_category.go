package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type CreateCategoryService struct {
	ctx context.Context
} // NewCreateCategoryService new CreateCategoryService
func NewCreateCategoryService(ctx context.Context) *CreateCategoryService {
	return &CreateCategoryService{ctx: ctx}
}

// Run create note info
func (s *CreateCategoryService) Run(req *product.CreateCategoryReq) (resp *product.CreateCategoryResp, err error) {
	category := &model.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	err = model.NewCategoryQuery(s.ctx, mysql.DB).CreateCategory(category)
	if err != nil {
		return
	}
	return &product.CreateCategoryResp{
		Id: uint32(category.ID),
	}, nil
}

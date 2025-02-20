package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// 初始化category查询对象
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	// 查询所有请求中的category
	categories := make([]model.Category, 0, len(req.Categories))
	for _, categoryName := range req.Categories {
		category, err := categoryQuery.GetByName(categoryName)
		if err != nil {
			return nil, err
		}
		if category == nil {
			continue
		}
		categories = append(categories, *category)
	}

	// 构建更新的商品对象
	updateProduct := &model.Product{
		Base: model.Base{
			ID: int(req.Id),
		},
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories,
	}

	// 更新商品
	productUpdate := model.NewProductUpdate(s.ctx, mysql.DB)
	if err := productUpdate.UpdateProduct(updateProduct); err != nil {
		return resp, nil
	}

	return
}

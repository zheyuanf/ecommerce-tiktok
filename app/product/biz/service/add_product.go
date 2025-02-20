package service

import (
	"context"
	"errors"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
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

	if len(categories) == 0 {
		return nil, errors.New("no category found")
	}

	// 创建新商品对象
	newProduct := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories, // 设置关联的categories
	}

	// 添加商品
	addProduct := model.NewAddProduct(s.ctx, mysql.DB)
	id, err := addProduct.AddProduct(newProduct)
	if err != nil {
		return nil, err
	}

	// 设置响应
	resp = &product.AddProductResp{
		Id: uint32(id),
	}
	return resp, nil
}

package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	product "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp = new(product.DeleteProductResp)
	// 创建删除服务实例
	productDelete := model.NewProductDelete(s.ctx, mysql.DB)
	// 执行删除操作
	err = productDelete.DeleteProduct(uint(req.Id))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

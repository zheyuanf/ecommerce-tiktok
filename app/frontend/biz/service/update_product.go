package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/infra/rpc"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	product "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/product"
	productrpc "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.UpdateProduct(h.Context, &productrpc.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Categories:  req.Categories,
	})
	return
}

package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	product "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/product"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	productrpc "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.AddProduct(h.Context, &productrpc.AddProductReq{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	return
}

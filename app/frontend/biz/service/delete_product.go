package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	product "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/product"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	productrpc "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.DeleteProduct(h.Context, &productrpc.DeleteProductReq{
		Id: req.Id,
	})
	return
}

package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	ctx := h.Context
	p, err := rpc.ProductClient.ListProducts(ctx, &product.ListProductsReq{})
	if err != nil {
		klog.Error(err)
	}
	var cartNum int
	return utils.H{
		"title":    "E706 E-commerce",
		"cart_num": cartNum,
		"items":    p.Products,
	}, nil
}

package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	category "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/category"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type CreateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateCategoryService(Context context.Context, RequestContext *app.RequestContext) *CreateCategoryService {
	return &CreateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateCategoryService) Run(req *category.CreateCategoryReq) (resp *common.Empty, err error) {
	_, err = rpc.ProductClient.CreateCategory(h.Context, &product.CreateCategoryReq{
		Name:        req.Name,
		Description: req.Description,
	})
	return
}

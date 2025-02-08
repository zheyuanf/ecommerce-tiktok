package cart

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/biz/service"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/biz/utils"
	cart "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/cart"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
)

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	_, err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusFound, []byte("/cart"))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// WarpResponse 补全返回给home模板的信息
func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	var cartNum int
	// 1. 获取user_id
	userId := frontendutils.GetUserIdFromCtx(ctx)

	// 2. 获取购物车数量
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: userId})
	if cartResp != nil && cartResp.Cart != nil {
		cartNum = len(cartResp.Cart.Items)
	}

	// 3. 补全返回给home模板的信息
	content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	content["cart_num"] = cartNum
	categoryResp, _ := rpc.ProductClient.ListCategories(ctx, &product.ListCategoriesReq{})
	content["categories"] = categoryResp.Categories
	return content
}

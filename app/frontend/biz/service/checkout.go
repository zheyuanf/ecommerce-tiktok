package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/checkout"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpccart "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	var items []map[string]string
	// 1. 获取用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	// 2. 获取购物车商品
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	// 3. 组装返回数据
	var total float32
	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}
	return utils.H{
		"title":    "Checkout",
		"items":    items,
		"cart_num": len(items),
		"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}

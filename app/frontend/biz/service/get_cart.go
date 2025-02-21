package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/utils"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpccart "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendutils.GetUserIdFromCtx(h.Context),
	})
	if err != nil {
		return nil, err
	}
	var items []map[string]string
	var total float64
	for _, item := range cartResp.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(item.Quantity) * float64(p.Price)
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}

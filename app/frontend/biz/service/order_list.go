package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/types"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpcorder "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	// 1. 获取用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	// 2. 调用RPC服务获取订单列表
	var orders []*types.Order
	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	if listOrderResp == nil || len(listOrderResp.Orders) == 0 {
		// 如果没有订单，则返回空列表
		return utils.H{
			"title":  "Order",
			"orders": orders,
		}, nil
	}
	// 3. 将订单列表转换为前端需要的格式
	for _, v := range listOrderResp.Orders {
		var items []types.OrderItem
		var total float32
		if len(v.OrderItems) > 0 {
			// 遍历订单项，获取商品信息并计算总价
			for _, vv := range v.OrderItems {
				total += vv.Cost
				i := vv.Item
				// 获取商品信息
				productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId})
				if err != nil {
					return nil, err
				}
				if productResp.Product == nil {
					continue
				}
				p := productResp.Product
				items = append(items, types.OrderItem{
					ProductId:   i.ProductId,
					Qty:         uint32(i.Quantity),
					ProductName: p.Name,
					Picture:     p.Picture,
					Cost:        vv.Cost,
				})
			}
		}
		timeObj := time.Unix(int64(v.CreatedAt), 0)
		// 组装订单信息并添加到订单列表中
		orders = append(orders, &types.Order{
			Cost:        total,
			Items:       items,
			CreatedDate: timeObj.Format("2006-01-02 15:04:05"),
			OrderId:     v.OrderId,
			Consignee:   types.Consignee{Email: v.Email},
		})
	}
	// 4. 返回订单列表
	return utils.H{
		"title":  "Order",
		"orders": orders,
	}, nil
}

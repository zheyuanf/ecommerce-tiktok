package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	order "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	// 1. 获取用户全部order
	orders, err := model.ListOrder(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	// 2. 获取每个order中的item
	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, v := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: v.Cost,
				Item: &cart.CartItem{
					ProductId: v.ProductId,
					Quantity:  v.Quantity,
				},
			})
		}
		o := &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				StreetAddress: v.Consignee.StreetAddress,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}
	return
}

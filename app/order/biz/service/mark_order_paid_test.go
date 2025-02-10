package service

import (
	"context"
	"testing"

	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	order "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
)

func TestMarkOrderPaid_Run(t *testing.T) {
	ctx := context.Background()

	mysql.MockInit()

	testCases := []struct {
		name     string
		placeReq *order.PlaceOrderReq
		markReq  *order.MarkOrderPaidReq
		err      error
	}{
		{
			name: "case 2",
			placeReq: &order.PlaceOrderReq{
				UserId:       123,
				UserCurrency: "USD",
				Address:      &order.Address{},
				Email:        "test@example.com",
				OrderItems: []*order.OrderItem{
					{
						Item: &cart.CartItem{
							ProductId: 321,
							Quantity:  2,
						},
						Cost: 100,
					},
				},
			},
			markReq: &order.MarkOrderPaidReq{
				UserId: 123,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			placeResp, _ := NewPlaceOrderService(ctx).Run(tc.placeReq)
			tc.markReq.OrderId = placeResp.Order.OrderId
			NewMarkOrderPaidService(ctx).Run(tc.markReq)
		})
	}

}

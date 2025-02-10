package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	order "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	mysql.MockInit()
	testCases := []struct {
		name string
		req  *order.PlaceOrderReq
		err  error
	}{
		{
			name: "case 1",
			req:  &order.PlaceOrderReq{},
			err:  fmt.Errorf("OrderItems empty"),
		},
		{
			name: "case 2",
			req: &order.PlaceOrderReq{
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
					{
						Item: &cart.CartItem{
							ProductId: 322,
							Quantity:  3,
						},
						Cost: 300,
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := s.Run(tc.req)
			assert.Equal(t, tc.err, err)
		})
	}
}

package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	cart "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

	tests := []struct {
		name    string
		req     *cart.AddItemReq
		wantErr bool
	}{
		{
			name: "normal case",
			req: &cart.AddItemReq{
				UserId: 1001,
				Item: &cart.CartItem{
					ProductId: 2001,
					Quantity:  2,
				},
			},
			wantErr: false,
		},
		{
			name: "zero quantity",
			req: &cart.AddItemReq{
				UserId: 1001,
				Item: &cart.CartItem{
					ProductId: 2001,
					Quantity:  0,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			s := NewAddItemService(ctx)

			resp, err := s.Run(tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

package main

import (
	"context"
	checkout "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/checkout"
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/biz/service"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}

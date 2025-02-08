package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/checkout"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpccheckout "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	// 1. 获取用户ID
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	// 2. 调用Checkout服务
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    userId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			City:          req.City,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
			CreditCardCvv:             req.Cvv,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}

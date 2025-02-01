package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	credit_card "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/biz/model"
	payment "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	// 1. 验证信用卡信息
	card := credit_card.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}
	// 2. 创建支付记录
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: transactionId.String()}, nil
}

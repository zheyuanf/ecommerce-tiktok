package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/infra/mq"
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/infra/rpc"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
	checkout "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/checkout"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/email"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/payment"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	// 1. 获取用户购物车数据
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		err = fmt.Errorf("GetCart.err:%v", err)
		return
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errors.New("cart is empty")
		return
	}
	// 2. 查询购物车中的每个商品，获取订单商品列表与总价格
	var (
		oi    []*order.OrderItem
		total float32 // 购物车总数
	)
	for _, cartItem := range cartResult.Cart.Items {
		// 获取商品
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	// 3. 创建订单
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		return
	}
	klog.Info("orderResult", orderResult)

	// 3. 清空购物车
	emptyResult, err := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		return
	}
	klog.Info(emptyResult)

	// 4. 付款
	orderId := orderResult.Order.OrderId
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	payResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}
	klog.Info(payResult)

	// 5. 邮件通知用户，向mq生产消息
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "abcde706@163.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "You just created an order in Ecommerce shop",
		Content:     "You just created an order in Ecommerce shop",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)

	// 6. 付款成功，修改订单状态
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: payResult.TransactionId,
	}
	return
}

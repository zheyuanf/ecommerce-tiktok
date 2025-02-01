package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/model"
	order "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.

	// 1. 判断订单是否为空
	if len(req.OrderItems) == 0 {
		err = fmt.Errorf("OrderItems empty")
		return
	}

	// 2. 创建订单与订单商品项
	mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 2.1 创建订单记录
		orderId, _ := uuid.NewUUID()
		o := &model.Order{
			OrderId:      orderId.String(),
			OrderState:   model.OrderStatePlaced,
			UserId:       req.UserId,
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email: req.Email,
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.Country = a.Country
			o.Consignee.State = a.State
			o.Consignee.City = a.City
			o.Consignee.StreetAddress = a.StreetAddress
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}

		// 2.2 创建订单内商品记录
		var itemList []*model.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &model.OrderItem{
				OrderIdRefer: o.OrderId,
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         v.Cost,
			})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}

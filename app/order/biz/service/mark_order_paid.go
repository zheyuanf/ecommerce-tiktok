package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/model"
	order "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	// 1. 校验用户id和订单id
	if req.UserId == 0 || req.OrderId == "" {
		err = fmt.Errorf("user_id or order_id can not be empty")
		return
	}
	// 2. 获取订单
	_, err = model.GetOrder(mysql.DB, s.ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	// 3. 更新订单状态为已支付
	err = model.UpdateOrderState(mysql.DB, s.ctx, req.UserId, req.OrderId, model.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return
}

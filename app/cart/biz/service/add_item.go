package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/model"
	cart "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(5000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}

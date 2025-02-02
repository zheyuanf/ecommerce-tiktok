package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/model"
	cart "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range list {
		items = append(items, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  v.Qty,
		})
	}

	c := &cart.Cart{
		Items:  items,
		UserId: req.UserId,
	}
	return &cart.GetCartResp{Cart: c}, nil
}

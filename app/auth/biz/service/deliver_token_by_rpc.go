package service

import (
	"context"
	"time"

	"github.com/zheyuanf/ecommerce-tiktok/app/auth/infra/token"
	auth "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 生成token，有效期为24小时
	token, err := token.TokenAuthenticator.GenerateToken(req.UserId, time.Unix(req.ExpireAt, 0))
	if err != nil {
		return nil, err
	}
	return &auth.DeliveryResp{
		Token: token,
	}, nil
}

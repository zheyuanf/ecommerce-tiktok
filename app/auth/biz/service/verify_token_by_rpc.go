package service

import (
	"context"

	"github.com/zheyuanf/ecommerce-tiktok/app/auth/infra/token"
	auth "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 解析token
	claims, err := token.TokenAuthenticator.ParseToken(req.GetToken())
	resp = &auth.VerifyResp{}
	if err == nil {
		resp.Res = true
		resp.UserId = claims.UserId
		resp.Role = claims.Role
	}
	return resp, nil
}

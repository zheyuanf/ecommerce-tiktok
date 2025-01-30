package service

import (
	"context"
	"testing"
	"time"

	"github.com/zheyuanf/ecommerce-tiktok/app/auth/infra/token"
	auth "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	token.InitAuthenticator("sign_key_test")

	req := &auth.DeliverTokenReq{
		UserId:   123,
		ExpireAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}

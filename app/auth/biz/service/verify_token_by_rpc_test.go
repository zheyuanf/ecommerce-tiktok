package service

import (
	"context"
	"testing"
	"time"

	"github.com/zheyuanf/ecommerce-tiktok/app/auth/infra/token"
	auth "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	ds := NewDeliverTokenByRPCService(ctx)
	token.InitAuthenticator("sign_key_test")
	userId := int32(123)
	dReq := &auth.DeliverTokenReq{
		UserId:   userId,
		ExpireAt: time.Now().Add(time.Second * 1).Unix(),
	}

	dResp, err := ds.Run(dReq)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", dResp)
	time.Sleep(time.Second * 2)

	vs := NewVerifyTokenByRPCService(ctx)
	vReq := &auth.VerifyTokenReq{
		Token: dResp.Token,
	}
	vResp, err := vs.Run(vReq)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", vResp)
}

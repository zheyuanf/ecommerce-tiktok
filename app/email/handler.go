package main

import (
	"context"
	email "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/email"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/biz/service"
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// Send implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) Send(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}

package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/auth"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpcuser "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	// 1. 调用 rpc 登录
	res, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return
	}

	// 2. 设置 session
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", res.UserId)
	err = session.Save()
	frontendutils.MustHandleError(err)

	// 3. 返回登录后的页面
	redirect := "/"
	if frontendutils.ValidateNext(req.Next) {
		// 如果存在 next 参数，则跳转到该页面
		redirect = req.Next
	}
	if err != nil {
		return "", err
	}

	return redirect, nil
}

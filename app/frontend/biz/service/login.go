package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	auth "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/auth"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	rpcauth "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
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
	userResp, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return
	}

	// 2. 分发token
	authResp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{
		UserId:   userResp.UserId,
		ExpireAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return
	}
	// TODO: 改成 localstorage
	// 3. 将 token 写入 cookie
	h.RequestContext.SetCookie("token", authResp.Token, int(time.Hour*24), "/", "",
		protocol.CookieSameSiteNoneMode, true, true)

	// 4. 返回登录后的页面
	redirect := "/"
	if frontendutils.ValidateNext(req.Next) {
		// 如果存在 next 参数，则跳转到该页面
		redirect = req.Next
	}

	return redirect, nil
}

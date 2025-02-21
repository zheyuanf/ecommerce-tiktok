package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	auth "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/auth"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/casbin"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	rpcuser "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	// 1. 注册用户
	userResp, err := rpc.UserClient.Register(h.Context, &rpcuser.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.Password,
		Role:            "user",
	})
	if err != nil {
		return nil, err
	}
	// 2. 添加用户的角色到casbin
	b, err := casbin.AddRoleForUser("user", strconv.Itoa(int(userResp.UserId)))
	if err != nil {
		return nil, err
	}
	if b == false {
		hlog.Info("用户与角色已经存在，无需重复添加")
	}
	return
}

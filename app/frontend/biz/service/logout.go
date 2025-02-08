package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	common "github.com/zheyuanf/ecommerce-tiktok/app/frontend/hertz_gen/frontend/common"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	// 清除session
	session := sessions.Default(h.RequestContext)
	session.Clear()
	session.Save() //nolint:errcheck
	return

}

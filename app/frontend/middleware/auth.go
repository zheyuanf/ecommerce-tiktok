package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/rpc"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth"
)

// GlobalAuth 保存 session 中的 userid 到 ctx
func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从 session 中获取 user_id
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId != nil {
			// 将 user_id 保存到 ctx 中，以便后续的 handler 可以使用
			ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		}
		c.Next(ctx)
	}
}

func GlobalJWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取 token
		// TODO: 暂时使用 cookie 方式获取 token
		token := string(c.Cookie("token"))
		req := &auth.VerifyTokenReq{
			Token: token,
		}
		// 调用 RPC 服务验证 token
		resp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, req)

		// 验证 token 是否有效，无效则重定向到登录页面
		if err == nil && resp.Res {
			// 将 user_id 保存到 ctx 中，以便后续的 handler 可以使用
			ctx = context.WithValue(ctx, utils.UserIdKey, resp.UserId)
		}
		c.Next(ctx)
	}
}

// Auth 鉴权中间件，没登录则跳转到 sign-in 界面
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		userId := session.Get("user_id")
		if userId == nil {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userId)
		c.Next(ctx)
	}
}

func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取 token
		token := string(c.Cookie("token"))
		req := &auth.VerifyTokenReq{
			Token: token,
		}
		// 调用 RPC 服务验证 token
		resp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, req)

		// 验证 token 是否有效，无效则重定向到登录页面
		if err != nil || !resp.Res {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = fmt.Sprintf("%s?next=%s", next, ref)
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
		// 将 user_id 保存到 ctx 中，以便后续的 handler 可以使用
		ctx = context.WithValue(ctx, utils.UserIdKey, resp.UserId)
		c.Next(ctx)
	}
}

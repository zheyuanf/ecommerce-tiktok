package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/casbin"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
)

func CasbinCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		role := ctx.Value(utils.UserRoleKey).(string)
		b, err := casbin.CasbinEnforcer.Enforce(role, string(c.Request.URI().Path()), string(c.Request.Method()))
		if err != nil || !b {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" {
				if utils.ValidateNext(ref) {
					next = ref
				}
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
	}
}

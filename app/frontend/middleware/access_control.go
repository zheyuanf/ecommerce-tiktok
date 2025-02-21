package middleware

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/casbin"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
)

func CasbinCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userId := ctx.Value(utils.UserIdKey).(int32)
		fmt.Println(strconv.Itoa(int(userId)), string(c.Request.URI().Path()), string(c.Request.Method()))
		b, err := casbin.CheckPermissionForUser(strconv.Itoa(int(userId)), string(c.Request.URI().Path()), string(c.Request.Method()))
		fmt.Println(b)
		if err != nil || !b {
			byteRef := c.GetHeader("Referer")
			ref := string(byteRef)
			next := "/sign-in"
			if ref != "" && utils.ValidateNext(ref) {
				next = ref
			}
			c.Redirect(302, []byte(next))
			c.Abort()
			c.Next(ctx)
			return
		}
	}
}

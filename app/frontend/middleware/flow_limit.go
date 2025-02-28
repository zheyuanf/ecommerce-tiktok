package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	zerolimit "github.com/zeromicro/go-zero/core/limit"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/infra/limit"
)

func FlowPeriodLimit(key string, f func(context.Context, *app.RequestContext) string) app.HandlerFunc {
	l, ok := limit.LimitMap[key].(*zerolimit.PeriodLimit)
	if !ok {
		hlog.Error(context.TODO(), "limit map error")
		return func(ctx context.Context, c *app.RequestContext) {
			c.String(500, "limit map error")
			c.Abort()
		}
	}
	return func(ctx context.Context, c *app.RequestContext) {
		s := f(ctx, c)
		val, err := l.Take(s)
		fmt.Println(val)
		if err != nil {
			hlog.Error(ctx, "limit take error")
			c.String(500, "limit take error")
			c.Abort()
			return
		}
		allow := true
		if val == zerolimit.OverQuota {
			allow = false
		}
		if allow {
			c.Next(ctx)
		} else {
			c.String(429, "too many requests")
			c.Abort()
		}
	}
}

func FlowTokenLimit(key string, f func(*app.RequestContext) string) app.HandlerFunc {
	l, ok := limit.LimitMap[key].(*zerolimit.TokenLimiter)
	if !ok {
		hlog.Error(context.TODO(), "limit map error")
		return func(ctx context.Context, c *app.RequestContext) {
			c.String(500, "limit map error")
			c.Abort()
		}
	}
	return func(ctx context.Context, c *app.RequestContext) {
		b := l.Allow()
		if b {
			c.Next(ctx)
		} else {
			c.String(429, "too many requests")
			c.Abort()
		}
	}
}

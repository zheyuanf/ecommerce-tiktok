package limit

import (
	zerolimit "github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/conf"
)

var (
	LimitMap map[string]any
)

func InitLimit() {
	LimitMap = make(map[string]any)
	// 用户登录等限流
	PutPeriodLimit(60, 10, "auth")
	// 商品搜索限流
	PutPeriodLimit(60, 100, "search_product")
	// 结算限流
	PutTokenLimit(60, 3, "checkout")
}

func PutPeriodLimit(period int, quota int, keyPrefix string) {
	l := zerolimit.NewPeriodLimit(period, quota, redis.New(conf.GetConf().Redis.Address), keyPrefix+"_")
	LimitMap[keyPrefix] = l
}

func PutTokenLimit(rate int, burst int, keyPrefix string) {
	l := zerolimit.NewTokenLimiter(rate, burst, redis.New(conf.GetConf().Redis.Address), keyPrefix+"_")
	LimitMap[keyPrefix] = l
}

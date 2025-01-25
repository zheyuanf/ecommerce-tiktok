package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

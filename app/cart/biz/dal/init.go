package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

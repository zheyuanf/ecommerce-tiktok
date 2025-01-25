package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

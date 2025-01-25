package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

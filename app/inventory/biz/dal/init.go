package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/auth/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

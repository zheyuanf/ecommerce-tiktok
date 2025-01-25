package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

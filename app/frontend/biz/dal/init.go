package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

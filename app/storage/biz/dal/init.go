package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/storage/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

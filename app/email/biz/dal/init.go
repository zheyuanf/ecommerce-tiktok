package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/email/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

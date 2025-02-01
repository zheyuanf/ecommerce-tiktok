package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/checkout/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

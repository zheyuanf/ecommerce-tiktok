package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

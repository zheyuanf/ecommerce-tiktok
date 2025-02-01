package dal

import (
	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

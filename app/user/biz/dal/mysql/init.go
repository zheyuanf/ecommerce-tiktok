package mysql

import (
	"fmt"
	"os"

	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	// 创建表结构
	if os.Getenv("GO_ENV") != "online" {
		err = DB.AutoMigrate(&model.User{})
	}
	if err != nil {
		panic(err)
	}
}

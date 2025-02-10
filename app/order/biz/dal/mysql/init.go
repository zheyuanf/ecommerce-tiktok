package mysql

import (
	"fmt"
	"os"

	"github.com/zheyuanf/ecommerce-tiktok/app/order/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/order/conf"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// 创建表结构
	if os.Getenv("GO_ENV") != "online" {
		err = DB.AutoMigrate(
			&model.Order{},
			&model.OrderItem{},
		)
	}
	if err != nil {
		panic(err)
	}
}

func MockInit() {
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(
		&model.Order{},
		&model.OrderItem{},
	)

	if err != nil {
		panic(err)
	}
}

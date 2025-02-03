package mysql

import (
	"fmt"
	"os"

	"github.com/zheyuanf/ecommerce-tiktok/app/product/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	err = DB.AutoMigrate( //nolint:errcheck
		&model.Product{},
		&model.Category{},
	)
	if err != nil {
		panic(err)
	}
}

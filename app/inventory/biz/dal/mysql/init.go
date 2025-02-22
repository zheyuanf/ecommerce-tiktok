package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/inventory/conf"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("Failed to open database connection: %v", err)
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// 自动同步表结构， 没有则创建表
	err = DB.AutoMigrate(&model.Inventory{})
	if err != nil {
		panic("failed to auto migrate")
	}

	log.Println("Successfully connected to MySQL database")
	return nil
}

func MockInit() {
	DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			// Logger:                 logger.Default.LogMode(logger.Silent),
			Logger: logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(
		&model.Inventory{},
	)

	if err != nil {
		panic(err)
	}
}

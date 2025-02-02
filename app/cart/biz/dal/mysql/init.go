package mysql

import (
	"fmt"
	"log"

	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() error {
	dsn := conf.GetConf().MySQL.DSN

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
	err = DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic("failed to auto migrate")
	}

	log.Println("Successfully connected to MySQL database")
	return nil
}

package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/model"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
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
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	// 自动同步表结构， 没有则创建表
	err = DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic("failed to auto migrate")
	}

	log.Println("Successfully connected to MySQL database")
	return nil
}

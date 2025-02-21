package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null"`
	Qty       int32  `gorm:"type:int(11);not null"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	var row Cart
	// 查找数据库中是否存在该商品
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 { // 已经存在该商品，数量+1
		return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty + ?", c.Qty)).Error
	}

	// 不存在创建
	return db.WithContext(ctx).Create(c).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id=?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var carts []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where("user_id=?", userId).Find(&carts).Error

	return carts, err
}

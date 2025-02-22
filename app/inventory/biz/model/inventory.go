package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID   int32 `gorm:"type:int(11);unique;not null;index:idx_product_id"` // 商品ID，外键
	Stock       int32 `gorm:"type:int(11);not null;default:0"`                   // 库存数量
	LockedStock int32 `gorm:"type:int(11);not null;default:0"`                   // 锁定库存数量
	Version     int32 `gorm:"type:int(11);not null;default:0"`                   // 版本号
}

func (Inventory) TableName() string {
	return "inventory"
}

// CreateInventory 创建库存记录
func CreateInventory(ctx context.Context, db *gorm.DB, inventory *Inventory) error {
	return db.WithContext(ctx).Create(inventory).Error
}

// DecreaseStock 扣减库存
func DecreaseStock(ctx context.Context, db *gorm.DB, productID int32, quantity int32) error {
	var inventory Inventory
	// 查询当前库存
	if err := db.WithContext(ctx).Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return err
	}
	// 检查库存是否足够
	if inventory.Stock-inventory.LockedStock < int32(quantity) {
		return fmt.Errorf("insufficient stock")
	}
	// 扣减库存
	inventory.Stock -= int32(quantity)

	// 更新版本号
	oldVerion := inventory.Version
	inventory.Version = oldVerion + 1

	// 更新库存记录
	ret := db.WithContext(ctx).
		Exec("UPDATE inventory SET stock = ?, version = ? WHERE product_id = ? AND version = ?",
			inventory.Stock, inventory.Version, productID, oldVerion)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected == 0 {
		return fmt.Errorf("optimistic lock error")
	}
	return nil
}

// IncreaseStock 增加库存
func IncreaseStock(ctx context.Context, db *gorm.DB, productID int32, quantity int32) error {
	inventory := Inventory{
		ProductID: productID,
	}
	err := db.WithContext(ctx).FirstOrCreate(&inventory, "product_id = ?", productID).Error
	if err != nil {
		return err
	}
	// 增加库存
	inventory.Stock += int32(quantity)

	// 更新版本号
	oldVerion := inventory.Version
	inventory.Version = oldVerion + 1

	// 更新库存记录
	ret := db.WithContext(ctx).Exec("UPDATE inventory SET stock = ?, version = ? WHERE product_id = ? AND version = ?",
		inventory.Stock, inventory.Version, productID, oldVerion)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected == 0 {
		return fmt.Errorf("optimistic lock error")
	}
	return nil
}

// 查询库存
func QueryInventory(ctx context.Context, db *gorm.DB, productID int) (*Inventory, error) {
	var inventory Inventory
	res := db.WithContext(ctx).Where("product_id = ?", productID).First(&inventory)
	if err := res.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &inventory, nil
}

// 批量查询库存
func BatchQueryInventory(ctx context.Context, db *gorm.DB, productIDs []int32) ([]*Inventory, error) {
	var inventories []*Inventory
	if err := db.WithContext(ctx).Where("product_id IN ?", productIDs).Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

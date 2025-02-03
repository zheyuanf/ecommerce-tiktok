package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func SearchProduct(db *gorm.DB, ctx context.Context, q string) (product []*Product, err error) {
	err = db.WithContext(ctx).Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return product, err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

func GetProductsByCategoryName(db *gorm.DB, ctx context.Context, name string) (category []Category, err error) {
	err = db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	return category, err
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p CategoryQuery) GetByName(categoryName string) (*Category, error) {
	var category Category
	result := p.db.WithContext(p.ctx).Where("name = ?", categoryName).First(&category)
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &category, err
}

func (p CategoryQuery) CreateCategory(category *Category) (err error) {
	result := p.db.WithContext(p.ctx).Model(&Category{}).Create(category)
	return result.Error
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) CategoryQuery {
	return CategoryQuery{ctx: ctx, db: db}
}

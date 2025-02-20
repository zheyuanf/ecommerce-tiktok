package model

import (
	"context"

	"github.com/redis/go-redis/v9"

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
	if err == gorm.ErrRecordNotFound {
		return Product{}, nil
	}
	return
}

func (p ProductQuery) SearchProduct(q string) (product []*Product, err error) {
	err = p.db.WithContext(p.ctx).
		Model(&Product{}).
		Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").
		Error
	return product, err
}

func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

type CachedProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	product, err = c.productQuery.GetById(productId)
	return
}

func (c CachedProductQuery) SearchProduct(q string) (product []*Product, err error) {
	return c.productQuery.SearchProduct(q)
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: NewProductQuery(ctx, db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

// 添加商品
type ProductInsert struct {
	ctx context.Context
	db  *gorm.DB
}

func (a ProductInsert) AddProduct(product *Product) (uint, error) {
	err := a.db.WithContext(a.ctx).Omit("Categories.*").Create(product).Error
	if err != nil {
		return 0, err
	}
	return uint(product.ID), nil
}

func NewAddProduct(ctx context.Context, db *gorm.DB) *ProductInsert {
	return &ProductInsert{
		ctx: ctx,
		db:  db,
	}
}

// 删除商品
type ProductDelete struct {
	ctx context.Context
	db  *gorm.DB
}

func (d ProductDelete) DeleteProduct(productId uint) error {
	d.db.WithContext(d.ctx).Select("Categories").Delete(&Product{}, productId)
	return nil
}

func NewProductDelete(ctx context.Context, db *gorm.DB) *ProductDelete {
	return &ProductDelete{
		ctx: ctx,
		db:  db,
	}
}

// 更新商品
type ProductUpdate struct {
	ctx context.Context
	db  *gorm.DB
}

func (u ProductUpdate) UpdateProduct(product *Product) error {
	return u.db.WithContext(u.ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除原有的商品分类关联
		if err := tx.Exec("DELETE FROM product_category WHERE product_id = ?", product.ID).Error; err != nil {
			return err
		}
		// 更新商品，并重新添加商品分类关联
		err := tx.Omit("Categories.*").Save(product).Error
		return err
	})
}

func NewProductUpdate(ctx context.Context, db *gorm.DB) *ProductUpdate {
	return &ProductUpdate{
		ctx: ctx,
		db:  db,
	}
}

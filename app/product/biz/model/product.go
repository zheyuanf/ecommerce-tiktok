package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`
	Storage     int32   `json:"storage"`

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
	// redis key
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	// search in redis
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)
	// error chain
	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cacheResultBytes, err := cachedResult.Bytes()
		if err != nil {
			return err
		}
		// redis中字符串转为product
		err = json.Unmarshal(cacheResultBytes, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	//从redis中读数据出错
	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		// 序列化为redis中存的字符串
		encode, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		// 存入redis， 过期时间1h
		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encode, time.Hour)
	}
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

// ProductMutation 数据库读写分离（暂时没用上）
type ProductMutation struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CachedProductQuery) CheckStorage(productId int, storage int32) (err error) {
	var product Product
	if err != nil {
		return
	}
	// redis分布式锁

	// redis key
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	//cachedKeyLock :=

	// get lock

	// search in redis
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)
	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cacheResultBytes, err := cachedResult.Bytes()
		if err != nil {
			return err
		}
		// redis中字符串转为product
		err = json.Unmarshal(cacheResultBytes, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	// redis中没有，直接更新DB
	if err != nil {
		// update DB
		product, err = c.productQuery.GetById(productId)
		err = c.productQuery.db.WithContext(c.productQuery.ctx).
			Model(&Product{}).Where("id=?", productId).Update("storage", product.Storage-storage).Error
		if err != nil {
			return
		}
	}
	product.Storage = product.Storage - storage
	encode, err := json.Marshal(product)
	if err != nil {
		return
	}
	_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encode, time.Hour)

	return
}

package model

import (
	"context"

	"gorm.io/gorm"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	Base
	Email          string `gorm:"unique"`
	PasswordHashed string
	Role           string
}

func (u User) TableName() string {
	return "user"
}

// 根据邮箱查询用户
func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

// 创建用户
func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

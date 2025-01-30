package service

import (
	"context"
	"errors"

	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/model"
	user "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	// 1. 参数校验
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}
	// 2. 获取user记录
	row, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return
	}

	// 3. 比对hash值是否正确
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return
	}
	return &user.LoginResp{UserId: int32(row.ID)}, nil
}

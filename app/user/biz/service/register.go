package service

import (
	"context"
	"errors"

	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/dal/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/user/biz/model"
	user "github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// 1. 参数校验
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		return nil, errors.New("email and password are required")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}
	if req.Role != model.RoleAdmin && req.Role != model.RoleUser {
		return nil, errors.New("role must be user or admin")
	}

	// 2. 计算密码的hash值
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// 3. 创建用户
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
		Role:           req.Role,
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}

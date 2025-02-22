package casbin

import (
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/conf"
)

var (
	CasbinEnforcer *casbin.Enforcer
)

func CasbinInit() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	a, err := gormadapter.NewAdapter("mysql", dsn)
	if err != nil {
		panic(err)
	}
	CasbinEnforcer, err = casbin.NewEnforcer("conf/rbac.conf", a)
	if err != nil {
		panic(err)
	}
}

// 添加权限
func AddPolicy(roleName, path, method string) bool {
	b, err := CasbinEnforcer.AddPolicy(roleName, path, method)
	if err != nil {
		return false
	}
	return b
}

// 添加角色下的用户
func AddRoleForUser(roleName string, userName string) (bool, error) {
	b, err := CasbinEnforcer.AddRoleForUser(userName, roleName)
	if err != nil {
		return false, err
	}
	return b, nil
}

// 检查权限
func CheckPermissionForRole(role, path, method string) (bool, error) {
	b, err := CasbinEnforcer.Enforce(role, path, method)
	return b, err
}

// 检查用户权限
func CheckPermissionForUser(userName, path, method string) (bool, error) {
	b, err := CasbinEnforcer.Enforce(userName, path, method)
	return b, err
}

package casbin

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAddRule(t *testing.T) {
	os.Chdir("../..")
	_ = godotenv.Load()
	CasbinInit()
	AddPolicy("user", "/checkout/result", "GET")
}

func TestAddRoleForUser(t *testing.T) {
	os.Chdir("../..")
	_ = godotenv.Load()
	CasbinInit()
	AddRoleForUser("user", "admin")
}

func TestCheckPermission(t *testing.T) {
	os.Chdir("../..")
	_ = godotenv.Load()
	CasbinInit()
	b, _ := CheckPermissionForRole("user", "/cart", "POST")
	t.Log(b)
}

func TestCheckPermissionForRole(t *testing.T) {
	os.Chdir("../..")
	_ = godotenv.Load()
	CasbinInit()
	b, _ := CheckPermissionForUser("3", "/cart", "GET")
	t.Log(b)
}

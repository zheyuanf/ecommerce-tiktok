package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/conf"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	"github.com/zheyuanf/ecommerce-tiktok/common/clientsuite"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/auth/authservice"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/order/orderservice"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/user/userservice"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	AuthClient     authservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: frontendutils.ServiceName,
		})
		initProductClient()
		initUserClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initAuthClient()
	})
}

func initProductClient() {
	// TODO: 熔断和降级等逻辑待补充
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	frontendutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	frontendutils.MustHandleError(err)
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", commonSuite)
	frontendutils.MustHandleError(err)
}

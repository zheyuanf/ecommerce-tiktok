package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/product"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/zheyuanf/ecommerce-tiktok/app/frontend/conf"
	frontendutils "github.com/zheyuanf/ecommerce-tiktok/app/frontend/utils"
	"github.com/zheyuanf/ecommerce-tiktok/common/clientsuite"
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
	})
}

func initProductClient() {
	// 熔断配置
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite,
		client.WithCircuitBreaker(cbs), client.WithFallback(
			fallback.NewFallbackPolicy(
				fallback.UnwrapHelper(
					func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
						if err == nil {
							return resp, nil
						}
						// 若为ListProducts方法，则返回一个默认的商品，防止界面为空
						methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
						if methodName != "ListProducts" {
							return resp, nil
						}
						return &product.ListProductsResp{
							Products: []*product.Product{
								{
									Price:       6.6,
									Id:          3,
									Picture:     "/static/image/t-shirt.jpeg",
									Name:        "T-shirt",
									Description: "static",
								},
							},
						}, nil
					}),
			),
		))
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

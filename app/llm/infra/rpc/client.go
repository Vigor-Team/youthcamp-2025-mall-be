package rpc

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/conf"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/clientsuite"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	UserClient    userservice.Client
	CartClient    cartservice.Client
	registryAddr  string
	commonSuite   client.Option
	serviceName   string
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		initProductClient()
		initOrderClient()
		initUserClient()
		initCartClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	utils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	utils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	utils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	utils.MustHandleError(err)
}

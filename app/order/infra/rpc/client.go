package rpc

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/order/conf"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/clientsuite"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
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
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	utils.MustHandleError(err)
}

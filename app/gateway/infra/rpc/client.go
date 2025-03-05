// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"
	"sync"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/retry"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/conf"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/infra/mtl"
	gatewayutils "github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/clientsuite"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	LlmClient      llmservice.Client
	PaymentClient  paymentservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonOpt      []client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonOpt = append(commonOpt, client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: gatewayutils.ServiceName,
		}))
		commonOpt = append(commonOpt, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))
		initProductClient()
		initUserClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
		initLlmClient()
		initPaymentClient()
	})
}

func initProductClient() {
	var opts []client.Option

	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("shop-gateway/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	opts = append(opts, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
		if err == nil {
			return resp, err
		}
		if methodName != "ListProducts" {
			return resp, err
		}
		return &product.ListProductsResp{
			Products: []*product.Product{
				{
					Price:       6.6,
					Id:          3,
					Picture:     "/static/image/t-shirt.jpeg",
					Name:        "T-Shirt",
					Description: "T-Shirt",
				},
			},
		}, nil
	}))))
	opts = append(opts, commonOpt...)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	gatewayutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonOpt...)
	gatewayutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonOpt...)
	gatewayutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonOpt...)
	gatewayutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonOpt...)
	gatewayutils.MustHandleError(err)
}

func initLlmClient() {
	var opts []client.Option
	fp := retry.NewFailurePolicy()
	fp.WithMaxRetryTimes(3)
	fp.WithMaxDurationMS(10000)
	fp.WithSpecifiedResultRetry(&retry.ShouldResultRetry{
		ErrorRetryWithCtx: func(ctx context.Context, err error, ri rpcinfo.RPCInfo) bool {
			if err == nil {
				return false
			}
			if kerrors.IsTimeoutError(err) {
				return true
			} else {
				return false
			}
		},
	})
	opts = append(opts, client.WithRetryMethodPolicies(map[string]retry.Policy{
		"SendMessage": {
			Enable:        true,
			Type:          0,
			FailurePolicy: fp,
		},
		"StreamMessage": {
			Enable:        true,
			Type:          0,
			FailurePolicy: fp,
		},
	}))
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("mall-gateway/llm/chat", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})
	opts = append(opts, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
		if err == nil {
			return resp, err
		}
		if methodName == "SendMessage" {
			return &llm.ChatResponse{
				Response: "服务器繁忙，请稍后再试",
			}, nil
		}
		if methodName == "StreamMessage" {
			return &llm.ChatResponse{
				Response: "服务器繁忙，请稍后再试",
			}, nil
		}
		return resp, err
	}))))
	LlmClient, err = llmservice.NewClient("llm", commonOpt...)
	gatewayutils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonOpt...)
	gatewayutils.MustHandleError(err)
}

package main

import (
	"context"
	"net"
	"strings"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal"
	dalmongo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/biz/dal/mongo"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/conf"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/llm/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/mtl"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/serversuite"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/utils"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/llm/llmservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"gopkg.in/natefinch/lumberjack.v2"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()

	dal.Init()

	rpc.InitClient()
	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	mtl.InitTracing(serviceName)
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	opts := kitexInit()
	opts = append(opts, server.WithGRPCInitialWindowSize(1024*1024*10), server.WithGRPCInitialConnWindowSize(1024*1024*10))

	svr := llmservice.NewServer(new(LlmServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	_ = provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)

	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))

	server.RegisterShutdownHook(func() {
		err := dalmongo.Client.Disconnect(context.TODO())
		if err != nil {
			klog.Error(err.Error())
			return
		}
	})
	return
}

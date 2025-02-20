package main

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/biz/dal"
	"github.com/zheyuanf/ecommerce-tiktok/app/payment/conf"
	"github.com/zheyuanf/ecommerce-tiktok/common/mtl"
	"github.com/zheyuanf/ecommerce-tiktok/common/serversuite"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/payment/paymentservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serviceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	MetricsPort  = conf.GetConf().Mtl.MetricsPort
	Endpoint     = conf.GetConf().Mtl.EndPoint
)

func main() {
	// 加载 .env 文件中的环境变量
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}
	mtl.InitMetric(serviceName, MetricsPort, RegistryAddr)
	p := mtl.InitTracing(serviceName, Endpoint)
	defer p.Shutdown(context.Background())
	// 初始化数据库连接
	dal.Init()

	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
	}))

	// 设置server的服务注册等配置
	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

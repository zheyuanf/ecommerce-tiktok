package main

import (
	"github.com/joho/godotenv"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/biz/dal"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/zheyuanf/ecommerce-tiktok/app/cart/conf"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/cart/cartservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	// 加载 .env 文件中的环境变量
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}

	// 初始化数据库连接
	dal.Init()

	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

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

	// 注册consul服务
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err.Error())
	}

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: serviceName,
	}), server.WithRegistry(r))

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

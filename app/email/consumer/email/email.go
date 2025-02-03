package email

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/infra/mq"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/infra/notify"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/email"
	"google.golang.org/protobuf/proto"
)

// 初始化消费者
func ConsumerInit() {
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req email.EmailReq
		// 从消息中解析req
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}
		// 发送邮件
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	// 注册服务关闭时，取消订阅，关闭nats连接
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe() //nolint:errcheck
		mq.Nc.Close()
	})
}

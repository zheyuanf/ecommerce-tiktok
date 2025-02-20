package email

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/infra/mq"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/infra/notify"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/email"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

// 初始化消费者
func ConsumerInit() {
	tracer := otel.Tracer("shop-nats-consumer")
	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req email.EmailReq
		// 从消息中解析req
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}
		// consumer otel
		ctx := context.Background()
		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(m.Header))
		_, span := tracer.Start(ctx, "shop-email-consumer")
		defer span.End()
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

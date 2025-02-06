package mq

import (
	"github.com/nats-io/nats.go"
	"github.com/zheyuanf/ecommerce-tiktok/app/email/conf"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(conf.GetConf().Nats.Address)
	if err != nil {
		panic(err)
	}
}

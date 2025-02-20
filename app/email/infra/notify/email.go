package notify

import (
	"crypto/tls"
	"github.com/zheyuanf/ecommerce-tiktok/rpc_gen/kitex_gen/email"
	"gopkg.in/gomail.v2"
	"log"
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	// 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", req.From)
	m.SetHeader("To", req.To)
	m.SetHeader("Subject", req.Subject)
	m.SetBody(req.ContentType, req.Content)
	d := gomail.NewDialer("smtp.163.com", 465, "abcde706@163.com", "JYafnACEvDNRiu9g")
	// 跳过验证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}

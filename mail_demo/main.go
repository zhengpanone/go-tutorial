package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	// 创建新的邮件消息
	m := gomail.NewMessage()
	// 设置邮件头部消息
	m.SetHeader("From", "zhengpanone@163.com")
	m.SetHeader("To", "zhengpanone@hotmail.com")
	m.SetHeader("Subject", "邮件标题")
	m.SetBody("text/html", "<h2>邮件内容</h2>")

	d := gomail.NewDialer(
		"smtp.163.com",
		25,
		"zhegnpanone@163.com",
		"xxxxx",
	)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（测试用）
	}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

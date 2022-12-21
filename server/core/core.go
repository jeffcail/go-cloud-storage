package core

import "os"

var (
	MailFrom       = os.Getenv("MailFrom")
	MailTo         = os.Getenv("MailTo")
	MailPassword   = os.Getenv("Mail163Pass")
	MailServer     = os.Getenv("MailServer")
	MailServerPort = os.Getenv("MailServerPort")
	MailTitle      = "cloud-storage 验证码发送测试"

	CodeLength  = 6
	CodeExpire  = 300
	TokenExpire = 3600

	QiNiuAK     = os.Getenv("QiNiuAK")
	QiNiuSk     = os.Getenv("QiNiuSk")
	QiNiuBucket = os.Getenv("QiNiuBucket")
	QiuNiuUrl   = os.Getenv("QiuNiuUrl")

	Page     = 1
	PageSize = 10
)

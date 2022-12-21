package test

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"testing"

	"github.com/jeffcail/cloud-storage/server/utils"

	"github.com/jeffcail/cloud-storage/server/core"

	"github.com/jordan-wright/email"
)

func TestRand(t *testing.T) {
	code := utils.RandCode()
	fmt.Println(code)
}

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = core.MailFrom
	e.To = []string{core.MailTo}
	e.Subject = "cloud-storage 验证码发送测试"
	e.HTML = []byte("您的验证码为:<h1>123456</h1>")
	err := e.SendWithTLS(fmt.Sprintf("%s%s", core.MailServer, core.MailServerPort),
		smtp.PlainAuth("", core.MailFrom,
			core.MailPassword, core.MailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: core.MailServer})
	if err != nil {
		t.Fatal(err)
	}
}

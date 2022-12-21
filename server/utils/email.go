package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jeffcail/cloud-storage/server/core"

	"github.com/jordan-wright/email"
)

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < core.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

// MailSendCode
func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = core.MailFrom
	e.To = []string{mail}
	e.Subject = core.MailTitle
	e.HTML = []byte("您的验证码为:<h1>" + code + "</h1>")
	err := e.SendWithTLS(fmt.Sprintf("%s%s", core.MailServer, core.MailServerPort),
		smtp.PlainAuth("", core.MailFrom, core.MailPassword, core.MailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: core.MailServer})
	if err != nil {
		return err
	}
	return nil
}

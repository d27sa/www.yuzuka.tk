package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
)

var (
	localMailAddr       = "yuzuka@gmx.com"
	localMailPassword   = "yuzuka.go"
	localMailHost       = "mail.gmx.com"
	localMailServerAddr = "mail.gmx.com:587"
	mailAuth            = smtp.PlainAuth("", localMailAddr, localMailPassword, localMailHost)
)

type mailer struct {
	addr       string
	serverAddr string
	auth       smtp.Auth
}

type mailers struct {
	a []*mailer
	i int
}

func newMailer(addr, pw, host string, port int) *mailer {
	return &mailer{
		addr,
		fmt.Sprintf("%s:%d", host, port),
		smtp.PlainAuth("", addr, pw, host),
	}
}

func (ms *mailers) get() *mailer {
	return ms.a[ms.i]
}

func (ms *mailers) next() *mailer {
	ms.i++
	return ms.a[ms.i]
}

func (ms *mailers) add(m *mailer) {
	ms.a = append(ms.a, m)
}

func newMailers() *mailers {
	return &mailers{
		make([]*mailer, 6),
		0,
	}
}

func sendVerificationMail(addr string) (string, error) {
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	to := []string{addr}
	msgStr := fmt.Sprintf("From: Yuzuka <admin@yuzuka.tk>\r\nTo: %s\r\nSender: <admin@yuzuka.tk>\r\nSubject: Verification code for your registeration\r\n\r\nHello!\r\n    Your verification code is %s.\r\n", addr, code)
	msg := []byte(msgStr)
	err := smtp.SendMail(localMailServerAddr, mailAuth, localMailAddr, to, msg)
	return code, err
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
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

func (m *mailer) sendVerificationMail(addr, code string) error {
	to := []string{addr}
	msgStr := fmt.Sprintf("From: Yuzuka <admin@yuzuka.tk>\r\nTo: %s\r\nSender: <admin@yuzuka.tk>\r\nSubject: Verification code for your registeration\r\n\r\nHello!\r\n    Your verification code is %s.\r\n", addr, code)
	msg := []byte(msgStr)
	err := smtp.SendMail(m.serverAddr, m.auth, m.addr, to, msg)
	return err
}

func (ms *mailers) next() *mailer {
	cur := ms.a[ms.i]
	ms.i++
	if ms.i == len(ms.a) {
		ms.i = 0
	}
	return cur
}

func (ms *mailers) num() int {
	return len(ms.a)
}

func (ms *mailers) add(m *mailer) {
	ms.a = append(ms.a, m)
}

func newMailers() *mailers {
	return &mailers{
		make([]*mailer, 0),
		0,
	}
}

func (ms *mailers) sendVerificationMail(addr string) (string, error) {
	var err error
	code := fmt.Sprintf("%04d", rand.Intn(10000))
	for i := 0; i < ms.num(); i++ {
		err = ms.next().sendVerificationMail(addr, code)
		if err == nil {
			return code, nil
		}
		log.Println(err)
	}
	return "", err
}

func initMailers() {
	serverMailers = newMailers()
	serverMailers.add(newMailer("kousuke@gmx.com", "kousuke.go", "mail.gmx.com", 587))
	serverMailers.add(newMailer("yuzuka@gmx.com", "yuzuka.go", "mail.gmx.com", 587))
	serverMailers.add(newMailer("shimotsuki@gmx.com", "kousuke.go", "mail.gmx.com", 587))
}

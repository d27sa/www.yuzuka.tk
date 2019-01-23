package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
)

var (
	localMailAddr       = "kousuke@gmx.com"
	localMailPassword   = "kousuke.go"
	localMailHost       = "mail.gmx.com"
	localMailServerAddr = "mail.gmx.com:587"
	mailAuth            = smtp.PlainAuth("", localMailAddr, localMailPassword, localMailHost)
)

func mailTest() {
	auth := smtp.PlainAuth("", localMailAddr, localMailPassword, localMailHost)
	to := []string{"erciyuangd@163.com"}
	msg := []byte("X-Sender: Kousuke\r\nTo: erciyuangd@163.com\r\n" +
		"Subject: こんにちは！\r\n" +
		"date: Wed, 23 Jan 2019 20:40:00 +0800\r\n" +
		"\r\n" +
		"今どこ？\r\n")
	err := smtp.SendMail("mail.gmx.com:587", auth, "kousuke@gmx.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendCheckMail(addr string) (int, error) {
	code := rand.Intn(10000)
	codeStr := strconv.Itoa(code)
	to := []string{addr}
	msgStr := fmt.Sprintf("From: Kousuke\r\nX-Sender: Kousuke\r\nTo: %s\r\nSubject: Check code from www.yuzuka.tk\r\n\r\nThe code is %s.\r\n", addr, codeStr)
	msg := []byte(msgStr)
	err := smtp.SendMail(localMailServerAddr, mailAuth, localMailAddr, to, msg)
	return code, err
}

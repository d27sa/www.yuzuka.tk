package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
)

func mailTest() {
	auth := smtp.PlainAuth("", "d27sa@yahoo.co.jp", "zhukm1997", "smtp.mail.yahoo.co.jp")
	to := []string{"erciyuangd@163.com"}
	msg := []byte("To: erciyuangd@163.com\r\n" +
		"Subject: こんにちは！\r\n" +
		"date: Wed, 23 Jan 2019 20:40:00 +0800\r\n" +
		"\r\n" +
		"今どこ？\r\n")
	err := smtp.SendMail("smtp.mail.yahoo.co.jp:587", auth, "d27sa@yahoo.co.jp", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func sendCheckMail(addr string) (int, error) {
	code := rand.Intn(10000)
	codeStr := strconv.Itoa(code)
	auth := smtp.PlainAuth("", "d27sa@yahoo.co.jp", "zhukm1997", "smtp.mail.yahoo.co.jp")
	to := []string{addr}
	msgStr := fmt.Sprintf("To: %s\r\nSubject: Check code from www.yuzuka.tk\r\n\r\nThe code is %s.\r\n", addr, codeStr)
	msg := []byte(msgStr)
	err := smtp.SendMail("smtp.mail.yahoo.co.jp:587", auth, "d27sa@yahoo.co.jp", to, msg)
	return code, err
}

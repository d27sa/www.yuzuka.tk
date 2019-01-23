package main

import (
	"log"
	"net/smtp"
)

func mailTest() {
	auth := smtp.PlainAuth("", "zkm971227@gmail.com", "zhukm19971227", "smtp.gmail.com")
	to := []string{"erciyuangd@163.com"}
	msg := []byte("To: erciyuangd@163.com\r\n" +
		"Subject: こんにちは！\r\n" +
		"date: Wed, 23 Jan 2019 20:40:00 +0800\r\n" +
		"\r\n" +
		"今どこ？\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "zkm971227@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

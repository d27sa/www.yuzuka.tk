package main

import (
	"log"
	"net/smtp"
)

func mailTest() {
	auth := smtp.PlainAuth("", "9d27sa@gmail.com", "zhukm1997", "smtp.gmail.com")
	to := []string{"d27sa@yahoo.co.jp"}
	msg := []byte("To: d27sa@yahoo.co.jp\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:465", auth, "9d27sa@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

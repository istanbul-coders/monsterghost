package main

import (
	"log"
	"net/smtp"
	"strconv"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func sendEmail(emailTo string, emailFrom string, subject string, body string, username string, password string) {
	to := []string{emailTo}
	msg := []byte("From: " + emailFrom + "\r\n" +
		"To: " + emailTo + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body +
		"\r\n")

	emailUser := &EmailUser{username, password, "smtp.gmail.com", 587}

	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

	err := smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port),
		auth,
		emailUser.Username,
		to,
		msg)
	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
	}

}

package main

import (
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

func sendEmail(emailTo string, emailCc string, emailFrom string, subject string, body string, username string, password string) {
	to := []string{emailTo}
	msg := []byte("From: " + emailFrom + "\r\n" +
	"To: " + emailTo + "\r\n" +
	"Cc: " + emailCc + "\r\n" +
	"Subject: " + subject + "\r\n" +
	"\r\n" +
	body +
	"\r\n")

	mailInformation := []string{emailTo, emailCc, emailFrom, subject, body}

	if containsByGivenCharacter("\\", mailInformation) {
		panic("You mustn't do escape character")
	}

	emailUser := &EmailUser{username, password, "smtp.gmail.com", 587}

	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

	err := smtp.SendMail(emailUser.EmailServer + ":" + strconv.Itoa(emailUser.Port),
		auth,
		emailUser.Username,
		to,
		msg)

	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
	}
}

func containsByGivenCharacter(char string, mailInformation []string) bool {
	for _, info := range mailInformation {
		if strings.ContainsAny(char, info) {
			return true
		}
	}
	return false
}

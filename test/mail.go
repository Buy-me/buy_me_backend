package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func main() {

	from := "binhdinhqt137@gmail.com"
	password := "rrvksahzyphwcicr"

	toEmailAddress := "binhdinhreact@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// subject := "Subject: This is the subject of the mail\n"
	// body := "This is the body of the mail"
	mail := Mail{
		Sender:  "Train Staff",
		To:      to,
		Subject: "Train Ticket",
		Body:    "Train Ticket Detail",
	}
	message := BuildMessage(mail)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, []byte(message))
	if err != nil {
		return
	}

	fmt.Println("Email sent successfully")
}

// func BuildMessage(name string) string {
// 	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
// 	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
// 	msg += fmt.Sprintf("To: %s\r\n", mail.To)
// 	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
// 	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
// 	return msg

// 	return msg
// }

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

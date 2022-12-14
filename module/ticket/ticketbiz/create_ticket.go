package ticketbiz

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
)

type CreateTicketStore interface {
	Create(context context.Context, data *ticketmodel.TicketCreate) error
}

type createTicketBiz struct {
	store CreateTicketStore
}

func NewCreateTicketBiz(store CreateTicketStore) *createTicketBiz {
	return &createTicketBiz{store: store}
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func (biz *createTicketBiz) CreateTicket(context context.Context, data *ticketmodel.TicketCreate) error {

	// from := "binhdinhqt137@gmail.com"
	// password := "rrvksahzyphwcicr"

	// toEmailAddress := data.Email
	// to := []string{toEmailAddress}

	// host := "smtp.gmail.com"
	// port := "587"
	// address := host + ":" + port

	// // subject := "Subject: This is the subject of the mail\n"
	// // body := "This is the body of the mail"

	// mail := Mail{
	// 	Sender:  "Super Train",
	// 	To:      to,
	// 	Subject: "Train Ticket",
	// 	Body:    "Train Ticket Detail",
	// }
	// message := BuildMessage(mail)

	// auth := smtp.PlainAuth("", from, password, host)

	// err := smtp.SendMail(address, auth, from, to, []byte(message))
	// if err != nil {
	// 	return err
	// }

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(ticketmodel.EntityName, err)
	}

	return nil
}

func BuildMessage(mail Mail) string {
	msg := ""
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

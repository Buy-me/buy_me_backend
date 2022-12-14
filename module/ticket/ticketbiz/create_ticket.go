package ticketbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/module/ticket/ticketmodel"
	"net/smtp"

	"github.com/matcornic/hermes/v2"
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

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(ticketmodel.EntityName, err)
	}

	go func() {
		from := "binhdinhqt137@gmail.com"
		password := "rrvksahzyphwcicr"

		toEmailAddress := data.Email
		to := []string{toEmailAddress}

		host := "smtp.gmail.com"
		port := "587"
		address := host + ":" + port

		// subject := "Subject: This is the subject of the mail\n"
		// body := "This is the body of the mail"

		auth := smtp.PlainAuth("", from, password, host)

		message := SendEmail()
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		err := smtp.SendMail(address, auth, from, to, []byte(mime+message))
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

// func BuildMessage(mail Mail) string {
// 	msg := ""
// 	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
// 	msg += fmt.Sprintf("To: %s\r\n", mail.To)
// 	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
// 	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

// 	return msg
// }

func SendEmail() string {
	h := hermes.Hermes{
		// Optional Theme

		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Thái Bình",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Công Vũ",
			Intros: []string{
				"Welcome to Super Train! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "This is your ticket detail",
					Button: hermes.Button{
						Color: "#f20", // Optional action button color
						Text:  "Table",
						Link:  "https://luv.vn/wp-content/uploads/2021/08/hinh-anh-gai-xinh-71.jpg",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	// emailText, err := h.GeneratePlainText(email)
	// if err != nil {
	// 	panic(err) // Tip: Handle error with something else than a panic ;)
	// }

	return emailBody
}

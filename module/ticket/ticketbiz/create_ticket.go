package ticketbiz

import (
	"context"
	"fmt"
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

type Model struct {
	Name   string `json:"name"`
	Millis int64  `json:"lastModified"`
}

func (biz *createTicketBiz) CreateTicket(context context.Context, data *ticketmodel.TicketCreate) error {

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(ticketmodel.EntityName, err)
	}

	go func(data *ticketmodel.TicketCreate) {
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

		message := SendEmail(data)
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		err := smtp.SendMail(address, auth, from, to, []byte(mime+message))
		if err != nil {
			panic(err)
		}
	}(data)
	return nil
}

func SendEmail(data *ticketmodel.TicketCreate) string {

	h := hermes.Hermes{
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			TroubleText: "Something went wrong",
			Name:        "Supper Train",
			Link:        "https://example-hermes.com/",
			Copyright:   "Copyright © 2020 Super Train. All rights reserved.",
			// Optional product logo
			Logo: "https://www.pngitem.com/pimgs/m/109-1099985_transparent-train-png-toy-train-clip-art-png.png",
		},
	}

	count := len(data.Travelers)
	// var data [][5]hermes.Entry

	table := make([][]hermes.Entry, count)
	table = append(table, []hermes.Entry{
		{Key: "Name", Value: "Name"},
		{Key: "Age", Value: "Age"},
		{Key: "Gender", Value: "Gender"},
		{Key: "Price", Value: "Price"},
		{Key: "National", Value: "National"},
	})
	for _, traveler := range data.Travelers {
		table = append(table, []hermes.Entry{
			{Key: "Name", Value: traveler.Name},
			{Key: "Age", Value: traveler.Age},
			{Key: "Gender", Value: traveler.Gender},
			{Key: "Price", Value: fmt.Sprintf("%f", traveler.Price)},
			{Key: "National", Value: traveler.National},
		})
	}

	dt, at := data.GetTime()
	email := hermes.Email{
		Body: hermes.Body{
			Name: data.Name,
			Intros: []string{
				"Welcome to Super Train! We're very excited to have you on board.\n Your Phone: " + data.Phone,
				"\n Departure Time: " + dt.Format("2006-01-02 15:04:05"),
				"\n Arrive Time: " + at.Format("2006-01-02 15:04:05"),
				"This is your ticket detail: ",
			},
			Table: hermes.Table{
				Data: table,
			},
			Outros: []string{
				fmt.Sprintf("Total Price: %f", data.TotalPrice),
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	return emailBody
}

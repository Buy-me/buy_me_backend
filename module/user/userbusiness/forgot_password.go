package userbusiness

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
	"log"
	"net/smtp"
)

type ForgotPasswordStore interface {
	Update(ctx context.Context, data usermodel.UserUpdate, id int) error
	FindUser(
		context context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type forgotPasswordBiz struct {
	store  ForgotPasswordStore
	hasher Hasher
}

func NewForgotPasswordBiz(store ForgotPasswordStore, hasher Hasher) *forgotPasswordBiz {
	return &forgotPasswordBiz{store: store, hasher: hasher}
}

func (biz *forgotPasswordBiz) ForgotPassword(context context.Context, email string) error {
	oldData, err := biz.store.
		FindUser(context, map[string]interface{}{"email": email})

	if err != nil {
		log.Println("Reset with email have not in System")
		return nil
	}

	var data usermodel.UserUpdate
	data.Password = "20222023"

	password := data.Password
	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	err = biz.store.Update(context, data, oldData.Id)

	if err != nil {
		return err
	}

	go func(email string, resetPassword string) {
		log.Println("email", email)
		log.Println("email", resetPassword)
		from := "binhdinhqt137@gmail.com"
		password := "rrvksahzyphwcicr"

		toEmailAddress := email
		to := []string{toEmailAddress}

		host := "smtp.gmail.com"
		port := "587"
		address := host + ":" + port

		subject := "Subject: This is the subject of the reset password\n"
		body := fmt.Sprintf("Your Reset Password: %s", resetPassword)
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		message := []byte(subject + mime + body)

		auth := smtp.PlainAuth("", from, password, host)

		err := smtp.SendMail(address, auth, from, to, message)
		if err != nil {
			panic(err)
		}
	}(email, password)

	return nil
}

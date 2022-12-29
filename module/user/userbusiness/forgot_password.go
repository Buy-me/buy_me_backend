package userbusiness

import (
	"context"
	"fmt"
	"food_delivery/common"
	"food_delivery/module/user/usermodel"
	"net/smtp"
)

type ForgotPasswordStore interface {
	Update(ctx context.Context, data usermodel.UserUpdate, id int) error
}

type forgotPasswordBiz struct {
	store  ForgotPasswordStore
	user   *usermodel.User
	hasher Hasher
}

func NewForgotPasswordBiz(store ForgotPasswordStore, user *usermodel.User, hasher Hasher) *forgotPasswordBiz {
	return &forgotPasswordBiz{store: store, user: user, hasher: hasher}
}

func (biz *forgotPasswordBiz) ForgotPassword(context context.Context, data usermodel.UserUpdate, id int) error {
	// oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	password := data.Password
	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	err := biz.store.Update(context, data, id)

	if err != nil {
		return err
	}

	go func(email string, resetPassword string) {
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
	}(biz.user.Email, password)

	return nil
}

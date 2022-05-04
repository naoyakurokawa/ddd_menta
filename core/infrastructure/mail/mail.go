package mail

import (
	gomail "gopkg.in/gomail.v2"
)

type GoMailer struct {
	Message *gomail.Message
}

func NewMailer() *GoMailer {
	return &GoMailer{
		Message: gomail.NewMessage(),
	}
}

// TODO 環境変数化、メールテンプレート化
func (mailer *GoMailer) Send(
	to string,
	subject string,
	body string,
) {
	mailer.Message.SetHeader("From", "test@example.com")
	mailer.Message.SetHeader("To", to)
	mailer.Message.SetHeader("Subject", subject)
	mailer.Message.SetBody("text/plain", body)

	d := gomail.Dialer{Host: "localhost", Port: 1025}
	if err := d.DialAndSend(mailer.Message); err != nil {
		panic(err)
	}
}

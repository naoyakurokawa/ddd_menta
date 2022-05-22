package mail

import (
	"github.com/naoyakurokawa/ddd_menta/config"
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

// TODO メールテンプレート化
func (mailer *GoMailer) Send(
	to string,
	subject string,
	body string,
) {
	mailer.Message.SetHeader("From", "test@example.com")
	mailer.Message.SetHeader("To", to)
	mailer.Message.SetHeader("Subject", subject)
	mailer.Message.SetBody("text/plain", body)

	d := gomail.Dialer{
		Host: config.Env.MailHost,
		Port: config.Env.MailPort,
	}
	if err := d.DialAndSend(mailer.Message); err != nil {
		panic(err)
	}
}

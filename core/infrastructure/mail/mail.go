package mail

import (
	"os"
	"strconv"

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

	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	d := gomail.Dialer{
		Host: host,
		Port: port,
	}
	if err := d.DialAndSend(mailer.Message); err != nil {
		panic(err)
	}
}

package mail

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	gomail "gopkg.in/gomail.v2"
)

type GoMailer struct {
	Message *gomail.Message
}

type Config struct {
	MailHost string `required:"true" default:"localhost""`
	MailPort int    `required:"true" default:"1025""`
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
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}
	mailer.Message.SetHeader("From", "test@example.com")
	mailer.Message.SetHeader("To", to)
	mailer.Message.SetHeader("Subject", subject)
	mailer.Message.SetBody("text/plain", body)

	d := gomail.Dialer{
		Host: config.MailHost,
		Port: config.MailPort,
	}
	if err := d.DialAndSend(mailer.Message); err != nil {
		panic(err)
	}
}

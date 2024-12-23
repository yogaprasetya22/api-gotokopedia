package mailer

import (
	"bytes"
	"text/template"
	"time"

	"github.com/go-mail/mail/v2"
)

type MailerConfig struct {
	Timeout      time.Duration
	Host         string
	Port         int
	Username     string
	Password     string
	Sender       string
	TemplatePath string
}

type Mailer struct {
	dailer *mail.Dialer
	config MailerConfig
	sender string
}

func New(config MailerConfig) Mailer {

	dailer := mail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	dailer.Timeout = config.Timeout

	return Mailer{
		dailer: dailer,
		sender: config.Sender,
		config: config,
	}
}

func (m Mailer) Send(templateFile, username, email string, data interface{}, isSendBox bool) error {

	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", "mochammad.yogaprasetya112@gmail.com")
	msg.SetHeader("Subject", subject.String())
	msg.SetHeader("From", m.sender)
	msg.AddAlternative("text/html", htmlBody.String())

	return m.dailer.DialAndSend(msg)
}

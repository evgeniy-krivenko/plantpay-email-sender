package mailer

import (
	"bytes"
	"embed"
	"github.com/go-mail/mail/v2"
	"html/template"
	"path"
	"time"
)

//go:embed "templates"
var templateFS embed.FS

type Config struct {
	Timeout      time.Duration
	Host         string
	Port         int
	Username     string
	Password     string
	Sender       string
	TemplatePath string
}

type Mailer struct {
	dialer *mail.Dialer
	sender string
	config Config
}

func New(c Config) *Mailer {
	dialer := mail.NewDialer(c.Host, c.Port, c.Username, c.Password)
	dialer.Timeout = c.Timeout
	return &Mailer{
		dialer: dialer,
		sender: c.Sender,
	}
}

func (m *Mailer) Send(to, subject, templateFile string, data interface{}) error {
	if m.config.TemplatePath == "" {
		m.config.TemplatePath = "templates/"
	}

	tmpl, err := template.New("email").ParseFS(templateFS, path.Join(m.config.TemplatePath, templateFile))
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetHeader("From", m.sender)
	msg.SetBody("text/html", htmlBody.String())
	return m.dialer.DialAndSend(msg)
}

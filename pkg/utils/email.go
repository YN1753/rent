package utils

import (
	"fmt"
	"rent/internal/config"

	"github.com/wneessen/go-mail"
)

func SendEmail(to string, subject string, body string) error {
	m := mail.NewMsg()
	fmt.Println(config.Cfg.Email)
	if err := m.From(config.Cfg.Email.Username); err != nil {
		return err
	}
	fmt.Println(to)
	if err := m.To(to); err != nil {
		return err
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)
	c, err := mail.NewClient(config.Cfg.Email.Host, mail.WithPort(config.Cfg.Email.Port), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(config.Cfg.Email.Username), mail.WithPassword(config.Cfg.Email.Password))
	if err != nil {
		return err
	}
	if err := c.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

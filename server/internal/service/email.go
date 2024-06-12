package service

import (
	"net/smtp"
)

type emailServiceImpl struct {
	auth smtp.Auth
	from string
	host string
}

func NewEmailService(username, password, from, host string) EmailService {
	auth := smtp.PlainAuth("", username, password, host)
	return &emailServiceImpl{auth: auth, from: from, host: host}
}

func (e *emailServiceImpl) SendMail(name, email, phone, msg string) error {
    subject := "Subject: New message from: " + email + "\n"
    body := name + "  |  " + email + "  |  " +  phone + "  |  " + msg
 
    err := smtp.SendMail(
        e.host+":587",
        e.auth,
        e.from,
        []string{e.from},
        []byte(subject+body),
    )
    return err
}



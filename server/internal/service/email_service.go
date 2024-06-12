package service

type EmailService interface {
	SendMail(name, email, phone, msg string) error
}

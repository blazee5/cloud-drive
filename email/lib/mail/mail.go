package mail

import (
	"net/smtp"
	"os"
)

func SendMail(email, code string) error {
	message := []byte("To: " + email + "\r\n" + "Subject: Ссылка для активации аккаунта\r\n" + "\r\n" + "Код: " + code + "\r\n")
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	err := smtp.SendMail(os.Getenv("SMTP_ADDR"), auth, os.Getenv("SMTP_FROM"), []string{email}, message)

	if err != nil {
		return err
	}

	return nil
}

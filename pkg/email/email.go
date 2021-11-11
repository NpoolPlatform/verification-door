package email

import "github.com/go-gomail/gomail"

const (
	Host     = "smtp.qq.com"
	Port     = 587
	Username = "crazyzplzpl@qq.com"
	Password = "xivtxurbagkwbbba"
)

func SendEmail(content string) error {
	dialer := gomail.NewDialer(Host, Port, Username, Password)

	message := gomail.NewMessage()
	message.SetAddressHeader("From", Username, Username)
	message.SetHeader("To", "crazyzplzpl@163.com")
	message.SetHeader("Subject", "Test")
	message.SetBody("text/html", content)

	return dialer.DialAndSend(message)
}

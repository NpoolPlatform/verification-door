package email

import (
	"context"
	"net/mail"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/mailgun/mailgun-go/v3"
	"golang.org/x/xerrors"
)

const (
	MailgunDomain = "mailgun_domain"
	MailgunApikey = "mailgun_apikey"
)

func SendEmail(content, email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return xerrors.Errorf("invalid email address: %v", err)
	}
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	domain := config.GetStringValueWithNameSpace(myServiceName, MailgunDomain)
	apikey := config.GetStringValueWithNameSpace(myServiceName, MailgunApikey)
	mg := mailgun.NewMailgun(domain, apikey)
	msg := mg.NewMessage(
		"Dear User <no_reply@"+domain+">",
		"test",
		content,
		email,
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()
	_, _, err = mg.Send(ctx, msg)
	return err
}

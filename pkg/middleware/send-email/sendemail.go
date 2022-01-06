package sendemail

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/email"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
)

const Sender = "email_sender"

func SendVerifyCode(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	var html, subtitle, user string
	switch in.GetLang() {
	case En:
		html = EnHTML
		subtitle = EnSubtitle
		user = EnUser
	case Jp:
		html = JpHTML
		subtitle = JpSubtitle
		user = JpUser
	default:
		html = JpHTML
		subtitle = JpSubtitle
		user = JpUser
	}

	var username string
	if in.GetUsername() == "" {
		username = user
	} else {
		if in.GetLang() == En {
			username = "Dear " + in.GetUsername() + ","
		} else {
			username += " 様、"
		}
	}

	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(ctx, in.GetEmail(), code, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	sender := config.GetStringValueWithNameSpace(myServiceName, Sender)

	err = email.SendEmailByAWS(subtitle, fmt.Sprintf(html, username, code), sender, in.GetEmail())
	if err != nil {
		return nil, err
	}

	return &npool.SendEmailResponse{
		Info: "send email successfully",
	}, nil
}

func SendUserSiteContactEmail(ctx context.Context, in *npool.SendUserSiteContactEmailRequest) (*npool.SendUserSiteContactEmailResponse, error) {
	text := fmt.Sprintf(SiteContactHTML, in.GetFrom(), in.GetUsername(), in.GetText())
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	sender := config.GetStringValueWithNameSpace(myServiceName, Sender)
	err := email.SendEmailByAWS(in.GetSubject(), text, sender, in.GetTo(), in.GetFrom())
	if err != nil {
		return nil, err
	}
	return &npool.SendUserSiteContactEmailResponse{
		Info: "send user site contact email successfully",
	}, nil
}

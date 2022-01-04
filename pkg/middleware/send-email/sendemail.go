package sendemail

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/email"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

const Sender = "email_sender"

func SendVerifyCode(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	var html, subtitle, user string
	switch in.Lang {
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

	if in.Username == "" {
		in.Username = user
	} else {
		if in.Lang == En {
			in.Username = "Dear " + in.Username + ","
		} else {
			in.Username += " 様、"
		}
	}

	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(ctx, in.Email, code, time.Now().Unix())
	if err != nil {
		return nil, xerrors.Errorf("fail to save email verify code: %v", err)
	}

	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	sender := config.GetStringValueWithNameSpace(myServiceName, Sender)

	err = email.SendEmailByAWS(subtitle, fmt.Sprintf(html, in.Username, code), sender, in.Email)
	if err != nil {
		return nil, xerrors.Errorf("fail to send email: %v", err)
	}

	return &npool.SendEmailResponse{
		Info: "send email successfully",
	}, nil
}

func SendUserSiteContactEmail(ctx context.Context, in *npool.SendUserSiteContactEmailRequest) (*npool.SendUserSiteContactEmailResponse, error) {
	err := email.SendEmailByAWS(in.SubTitle, in.Text, in.From, in.To)
	if err != nil {
		return nil, err
	}
	return &npool.SendUserSiteContactEmailResponse{
		Info: "send user site contact email successfully",
	}, nil
}

package sendemail

import (
	"context"
	"time"

	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/email"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

func SendEmail(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	switch in.Intention {
	case "verify":
		return VerifyCode(ctx, in)
	case "":
		return VerifyCode(ctx, in)
	}
	return nil, xerrors.Errorf("you need to choose a correct send reason.")
}

func VerifyCode(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(in.Email, code, time.Now().Unix())
	if err != nil {
		return nil, xerrors.Errorf("fail to save email verify code: %v", err)
	}

	err = email.SendEmail(`
	<p>your email code is: </p>
	`+code, in.Email)
	if err != nil {
		return nil, xerrors.Errorf("fail to send email: %v", err)
	}

	return &npool.SendEmailResponse{
		Info: "send email successfully",
	}, nil
}

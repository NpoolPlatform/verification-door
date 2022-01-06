package sendsms

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/sms"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

const (
	En        = "en-US"
	Jp        = "ja-JP"
	EnMessage = "[Procyon] Verification Code: %v\nThis code is valid for 10 minutes\nPlease do not share this code with anyone."
	JpMessage = "[Procyon] 認証コード: %v\n利用期間: 10分\nこのコードを他人と共有するのはお控えください。"
)

func SendVerifyCode(ctx context.Context, in *npool.SendSmsRequest) (*npool.SendSmsResponse, error) {
	if in.GetPhone() == "" {
		return nil, xerrors.Errorf("please input your phone number")
	}
	var message string

	switch in.GetLang() {
	case En:
		message = EnMessage
	default:
		message = JpMessage
	}

	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(ctx, in.GetPhone(), code, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	err = sms.SendSms(fmt.Sprintf(message, code), in.GetPhone())
	if err != nil {
		return nil, xerrors.Errorf("fail to send sms message: %v", err)
	}

	return &npool.SendSmsResponse{
		Info: "send sms message successfully",
	}, nil
}

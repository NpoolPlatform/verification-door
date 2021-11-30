package sendemail

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/email"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

const (
	En         = "en_US"
	Jp         = "ja_JP"
	Ch         = "zh_CN"
	From       = "Procyon <support@procyon.vip>"
	JpSubtitle = "【Procyon】Email認証コード"
	EnSubtitle = "【Procyon】Email verification code"
	ChSubtitle = "【Procyon】邮件验证码"
	JpHTML     = `
	<html>
  <head>

  </head>

  <body>
    <div>
			<img src = 'https://s3.bmp.ovh/imgs/2021/11/d628183d307ac487.jpg' />
      <p>Procyonユーザーへ</p>
      <p>日頃より弊社サービスをご愛顧頂き誠にありがとうございます。</p>
      <h3>Email認証コード: <span style="color: #1ec498;">%v</span></h3>
      <p>ご不明な点がございましたら、カスタマーサポートへまでお問い合わせください。</p>
      <p>今後ともどうぞ宜しくお願い申し上げます。</p>
      <p>----------------</p>
      <p>Procyon カスタマーサポートセンター</p>
      <p>Email: <a href="mailto:support@procyon.vip">support@procyon.vip</a></p>
      <p>営業時間: 10時00分 〜 22時00分 (日本時間)</p>
    </div>
  </body>
</html>
	`
	EnHTML = `
	<html>
  <head>

  </head>

  <body>
    <div>
			<img src = 'https://s3.bmp.ovh/imgs/2021/11/d628183d307ac487.jpg' />
      <p>To Procyon users</p>
      <p>Thank you very much for your continued patronage of our services.</p>
      <h3>Email Verification Code: <span style="color: #1ec498;">%v</span></h3>
      <p>If you have any questions, please contact our customer support.</p>
      <p>We are looking forward to working with you in the future.</p>
      <p>----------------</p>
      <p>Procyon Customer Support Center</p>
      <p>Email: <a href="mailto:support@procyon.vip">support@procyon.vip</a></p>
      <p>Business hours: 10:00 am - 10:00 pm (Japan time)</p>
    </div>
  </body>
</html>
	`
	ChHTML = `
	<html>
  <head>

  </head>

  <body>
    <div>
			<img src = 'https://s3.bmp.ovh/imgs/2021/11/d628183d307ac487.jpg' />
      <p>Procyon的用户</p>
      <p>非常感谢您对我们的服务感兴趣。</p>
      <h3>Email验证码: <span style="color: #1ec498;">%v</span></h3>
      <p>如果你有任何问题，请联系我们的客户支持团队。</p>
      <p>我们期待着在未来与您合作。</p>
      <p>----------------</p>
      <p>Procyon客户支持中心</p>
      <p>Email: <a href="mailto:support@procyon.vip">support@procyon.vip</a></p>
      <p>营业时间：上午10点至晚上10点（GMT）</p>
    </div>
  </body>
</html>
	`
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
	var content, subtitle string
	switch in.Lang {
	case Ch:
		content = ChHTML
		subtitle = ChSubtitle
	case En:
		content = EnHTML
		subtitle = EnSubtitle
	case Jp:
		content = JpHTML
		subtitle = JpSubtitle
	default:
		return nil, xerrors.Errorf("input lang <%v> doesn't exist!!!", in.Lang)
	}

	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(in.Email, code, time.Now().Unix())
	if err != nil {
		return nil, xerrors.Errorf("fail to save email verify code: %v", err)
	}

	err = email.SendEmail(From, subtitle, fmt.Sprintf(content, code), in.Email)
	if err != nil {
		return nil, xerrors.Errorf("fail to send email: %v", err)
	}

	return &npool.SendEmailResponse{
		Info: "send email successfully",
	}, nil
}

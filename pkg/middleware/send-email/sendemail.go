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
	En         = "en-US"
	Jp         = "ja-JP"
	Ch         = "zh-CN"
	JpSubtitle = "【Procyon】Email認証コード"
	EnSubtitle = "【Procyon】Email verification code"
	ChSubtitle = "【Procyon】邮件验证码"
	JpUser     = "Procyonユーザーへ"
	EnUser     = "Dear Procyon User,"
	CnUser     = "Procyon的用户"
	JpHTML     = `
	<html>
  <head>

  </head>

  <body>
    <div>
			<img src = 'https://s3.bmp.ovh/imgs/2021/12/6d8fbb9f12f2b978.jpg' />
      <p>%v</p>
      <p>日頃より弊社サービスをご愛顧頂き誠にありがとうございます。</p>
      <p>以下、認証コードをログインの際にご利用ください。</p>
      <h3>Email認証コード: <span style="color: #1ec498;">%v</span></h3>
      <p>ご不明な点がございましたら、カスタマーサポートへまでお問い合わせください。</p>
      <p>今後ともどうぞ宜しくお願い申し上げます。</p>
      <p>----------------</p>
      <p></p>
    </div>

		<div>
      <table style="padding: 0px; margin: 10px 0; border: none">
        <tbody>
          <tr>
            <td style="vertical-align: middle; padding: 10px 7px 0 0">
              <img
                alt="Img"
                src="https://storage.googleapis.com/signaturesatori/customer-C01934n52/images/companyLogo/IITLT.png"
                style="border: 0px"
                width="80"
                height="80"
              />
            </td>
            <td
              style="
                border-left: 3px solid #b3bdd0;
                padding: 7px 0px 0px 10px;
                font-family: 'calibri', 'candara', 'segoe', 'segoe ui', 'optima',
                  'arial', sans-serif;
                font-size: 12px;
                line-height: 14px;
                color: #3a3a41;
              "
            >
              <div style="margin-bottom: 10px">
                <strong>
                  <span style="color: #3a3a41">
                    <span style="font-size: 14px; line-height: 17px"
                      >Procyon カスタマーサポートセンター</span
                    >
                  </span>
                </strong>
                <br />
              </div>
              <div style="margin-bottom: 10px">
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="mailto:support@procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >support@procyon.vip</a
                  >
                </span>
              </div>
              <div style="margin-bottom: 10px">
                <strong>Procyon</strong>
                <br />
                営業時間: 10時00分 〜 22時00分 (日本時間)
                <br />
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="https://www.procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >www.procyon.vip</a
                  >
                </span>
              </div>
              <div></div>
            </td>
          </tr>
        </tbody>
      </table>
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
			<img src = 'https://s3.bmp.ovh/imgs/2021/12/6d8fbb9f12f2b978.jpg' />
      <p>%v</p>
      <p>Thank you very much for using our services.</p>
      <p>Please use this code to log in</p>
      <h3>Email Verification Code: <span style="color: #1ec498;">%v</span></h3>
      <p>If you have any questions about logging in, please reply to this email.</p>
      <p>Respectfully, we are looking forward to working with you in the future.</p>
      <p>----------------</p>
      <p></p>
    </div>

		<div>
      <table style="padding: 0px; margin: 10px 0; border: none">
        <tbody>
          <tr>
            <td style="vertical-align: middle; padding: 10px 7px 0 0">
              <img
                alt="Img"
                src="https://storage.googleapis.com/signaturesatori/customer-C01934n52/images/companyLogo/IITLT.png"
                style="border: 0px"
                width="80"
                height="80"
              />
            </td>
            <td
              style="
                border-left: 3px solid #b3bdd0;
                padding: 7px 0px 0px 10px;
                font-family: 'calibri', 'candara', 'segoe', 'segoe ui', 'optima',
                  'arial', sans-serif;
                font-size: 12px;
                line-height: 14px;
                color: #3a3a41;
              "
            >
              <div style="margin-bottom: 10px">
                <strong>
                  <span style="color: #3a3a41">
                    <span style="font-size: 14px; line-height: 17px"
                      >Procyon Customer Support Center</span
                    >
                  </span>
                </strong>
                <br />
              </div>
              <div style="margin-bottom: 10px">
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="mailto:support@procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >support@procyon.vip</a
                  >
                </span>
              </div>
              <div style="margin-bottom: 10px">
                <strong>Procyon</strong>
                <br />
                Business hours: 10:00 am - 10:00 pm (Japan time)
                <br />
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="https://www.procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >www.procyon.vip</a
                  >
                </span>
              </div>
              <div></div>
            </td>
          </tr>
        </tbody>
      </table>
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
      <p>%v</p>
      <p>非常感谢您对我们的服务感兴趣。</p>
      <h3>Email验证码: <span style="color: #1ec498;">%v</span></h3>
      <p>如果你有任何问题，请联系我们的客户支持团队。</p>
      <p>我们期待着在未来与您合作。</p>
      <p>----------------</p>
      <p></p>
    </div>

		<div>
      <table style="padding: 0px; margin: 10px 0; border: none">
        <tbody>
          <tr>
            <td style="vertical-align: middle; padding: 10px 7px 0 0">
              <img
                alt="Img"
                src="https://storage.googleapis.com/signaturesatori/customer-C01934n52/images/companyLogo/IITLT.png"
                style="border: 0px"
                width="80"
                height="80"
              />
            </td>
            <td
              style="
                border-left: 3px solid #b3bdd0;
                padding: 7px 0px 0px 10px;
                font-family: 'calibri', 'candara', 'segoe', 'segoe ui', 'optima',
                  'arial', sans-serif;
                font-size: 12px;
                line-height: 14px;
                color: #3a3a41;
              "
            >
              <div style="margin-bottom: 10px">
                <strong>
                  <span style="color: #3a3a41">
                    <span style="font-size: 14px; line-height: 17px"
                      >Procyon客户支持中心</span
                    >
                  </span>
                </strong>
                <br />
              </div>
              <div style="margin-bottom: 10px">
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="mailto:support@procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >support@procyon.vip</a
                  >
                </span>
              </div>
              <div style="margin-bottom: 10px">
                <strong>Procyon</strong>
                <br />
                营业时间：上午10点至晚上10点（GMT）
                <br />
                <span style="color: #3a3a41; text-decoration: none">
                  <a
                    href="https://www.procyon.vip"
                    style="color: #3a3a41; text-decoration: none"
                    target="_blank"
                    >www.procyon.vip</a
                  >
                </span>
              </div>
              <div></div>
            </td>
          </tr>
        </tbody>
      </table>
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
	var html, subtitle, user string
	switch in.Lang {
	case Ch:
		html = ChHTML
		subtitle = ChSubtitle
		user = CnUser
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
		in.Username = "Dear " + in.Username
	}

	code := verifycode.GenerateVerifyCode(6)
	err := verifycode.SaveVerifyCode(in.Email, code, time.Now().Unix())
	if err != nil {
		return nil, xerrors.Errorf("fail to save email verify code: %v", err)
	}

	err = email.SendEmailByAWS(subtitle, fmt.Sprintf(html, in.Username, code), in.Email)
	if err != nil {
		return nil, xerrors.Errorf("fail to send email: %v", err)
	}

	return &npool.SendEmailResponse{
		Info: "send email successfully",
	}, nil
}

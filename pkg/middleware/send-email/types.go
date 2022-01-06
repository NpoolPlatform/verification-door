package sendemail

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
	SiteContactHTML = `
  <html>
  <head>

  </head>

  <body>
    <div>From: %v, username is: %v</div>
    <div>%v</div>
  </body>
  </html>
  `
)

package captcher

import (
	"strings"

	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/xerrors"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcherImg() (*npool.GetCaptcherImgResponse, error) {
	driver := &base64Captcha.DriverDigit{
		Height:   80,
		Width:    200,
		Length:   5,
		MaxSkew:  0.7,
		DotCount: 80,
	}

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	return &npool.GetCaptcherImgResponse{
		Info: &npool.CaptcherResp{
			ID:        id,
			Base64URL: b64s,
		},
	}, nil
}

func VerifyCaptcher(in *npool.VerifyCaptcherRequest) error {
	pass := store.Verify(in.ID, strings.TrimSpace(in.Input), true)
	if !pass {
		return xerrors.Errorf("input code not correct")
	}
	return nil
}

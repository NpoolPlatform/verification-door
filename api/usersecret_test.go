package api

import (
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/verification"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAuthAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	userID := uuid.New().String()
	username := uuid.New().String()
	appID := uuid.New().String()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetQRcodeURLRequest{
			Username: username,
			UserID:   userID,
			AppID:    appID,
		}).Post("http://localhost:50090/v1/get/qrcode/url")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.VerifyGoogleAuthRequest{
			UserID: userID,
			Code:   "11111",
			AppID:  appID,
		}).Post("http://localhost:50090/v1/verify/google/auth")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp1.StatusCode())
	}

	resp2, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.DeleteUserGoogleAuthRequest{
			UserID: userID,
			AppID:  appID,
		}).Post("http://localhost:50090/v1/delete/user/google/auth")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp2.StatusCode())
	}
}

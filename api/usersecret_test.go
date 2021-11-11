package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
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

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.GetQRcodeURLRequest{
			Username: username,
			UserID:   userID,
		}).Post("http://localhost:32759/v1/get/qrcode/url")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	_, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.VerifyGoogleAuthRequest{
			UserID: userID,
			Code:   "11111",
		}).Post("http://localhost:32759/v1/verify/google/auth")
	assert.Nil(t, err)

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.DeleteUserGoogleAuthRequest{
			UserID: userID,
		}).Post("http://localhost:32759/v1/delete/user/google/auth")
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
	}
}

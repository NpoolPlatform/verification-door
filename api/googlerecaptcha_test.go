package api

import (
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/verification"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGoogleRecaptchaAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&npool.VerifyGoogleRecaptchaRequest{
			Response: "test",
		}).Post("http://localhost:50090/v1/verify/google/recaptcha")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp.StatusCode())
	}
}

package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
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
	fmt.Println(resp, "err is", err)
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp.StatusCode())
	}
}

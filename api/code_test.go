package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestVerifyCodeAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	userID := uuid.New().String()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SendEmailRequest{
			UserID: userID,
			Email:  "crazyzplzpl@163.com",
		}).Post("http://localhost:32759/v1/send/email")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
	}

	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.VerifyCodeRequest{
			UserID:     userID,
			Code:       "111111",
			VerifyType: "email",
		}).Post("http://localhost:32759/v1/verify/code")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp1.StatusCode())
	}
}

package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/verification"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		logger.Sugar().Error(err)
	}
}

func TestVerifyCodeAPI(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	resp1, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.VerifyCodeRequest{
			Param:      "crazyzplzpl@163.com",
			Code:       "111111",
			VerifyType: "email",
		}).Post("http://localhost:50090/v1/verify/code")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp1.StatusCode())
	}
}

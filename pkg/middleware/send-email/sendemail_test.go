package sendemail

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
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

func TestSendEmailMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	email := "crazyzplzpl@163.com"
	resp, err := SendEmail(context.Background(), &npool.SendEmailRequest{
		Intention: "verify",
		Email:     email,
		Lang:      Jp,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	_, err = VerifyCode(context.Background(), &npool.SendEmailRequest{
		Email: email,
	})
	assert.NotNil(t, err)
}

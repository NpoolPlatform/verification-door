package verifycode

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func TestVerifyCode(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	code := GenerateVerifyCode(6)
	assert.NotNil(t, code)

	userEmail := "crazyzplzpl@qq.com"
	sendTime := time.Now().Unix()
	err := SaveVerifyCode(context.Background(), userEmail, code, sendTime)
	assert.Nil(t, err)

	_, err = VerifyCode(context.Background(), &npool.VerifyCodeRequest{
		Code:  code,
		Param: userEmail,
	})
	assert.Nil(t, err)
}

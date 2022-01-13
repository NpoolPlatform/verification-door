package verifycode

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/verification"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		log.Fatal(err)
	}
}

func TestVerify(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	_, err := VerifyCodeWithUserID(context.Background(), &npool.VerifyCodeWithUserIDRequest{
		UserID: "fcba02d3-c98e-47d1-b98f-879d5f8e005c",
		Param:  "crazyzplzpl@163.com",
		Code:   "123456",
	})
	assert.NotNil(t, err)
}

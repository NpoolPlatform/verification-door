package verifycode

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestVerify(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	resp, err := VerifyCodeWithUserID(&npool.VerifyCodeWithUserIDRequest{
		UserID: "fcba02d3-c98e-47d1-b98f-879d5f8e005c",
		Param:  "crazyzplzpl@163.com",
		Code:   "123456",
	})
	fmt.Println(resp, err)
}

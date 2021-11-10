package verifycode

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
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

func TestVerifyCode(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	code := GenerateVerifyCode(6)
	assert.NotNil(t, code)

	userID := uuid.New().String()
	sendType := "email"
	sendTime := time.Now().Unix()
	err := SaveVerifyCode(userID, code, sendType, sendTime)
	assert.Nil(t, err)

	err = VerifyCode(userID, code, sendType)
	assert.Nil(t, err)
}

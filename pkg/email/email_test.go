package email

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
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

func TestEmail(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := SendEmail("crazyzplzpl@qq.com", "test", "test", "", "crazyzplpzl@163.com")
	assert.Nil(t, err)
	err = SendEmailByAWS("test", "test", "crazyzplzpl@gmail.com")
	assert.Nil(t, err)
}

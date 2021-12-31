package google

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/google/uuid"
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

func TestGoogle(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	secret, err := GetSecret()
	if assert.Nil(t, err) {
		assert.NotNil(t, secret)
	}

	userID := uuid.New().String()
	url := GetQrcodeURL(userID, secret)
	assert.NotNil(t, url)

	trueCodes, err := getTrueCode(secret)
	if assert.Nil(t, err) {
		assert.NotNil(t, trueCodes)
	}

	result, err := VerifyCode(secret, trueCodes[0])
	if assert.Nil(t, err) {
		assert.Equal(t, result, true)
	}
}

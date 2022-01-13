package googleauth

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/verification"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/google/uuid"
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

func TestGoogleAuthMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	username := "test"
	userID := uuid.New().String()
	appID := uuid.New().String()
	resp, err := GetQRcodeURL(context.Background(), &npool.GetQRcodeURLRequest{
		Username: username,
		UserID:   userID,
		AppID:    appID,
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	_, err = VerifyGoogleAuth(context.Background(), &npool.VerifyGoogleAuthRequest{
		UserID: userID,
		Code:   "12345",
		AppID:  appID,
	})
	assert.NotNil(t, err)
}

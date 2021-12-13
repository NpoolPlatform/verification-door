package googleauth

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
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
	fmt.Println("resp is ", resp.Info)
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}
	fmt.Println(userID)

	_, err = VerifyGoogleAuth(context.Background(), &npool.VerifyGoogleAuthRequest{
		UserID: userID,
		Code:   "12345",
		AppID:  appID,
	})
	assert.NotNil(t, err)
}

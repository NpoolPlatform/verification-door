package usersecret

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

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

func TestUserSecretCRUD(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	secret := "123456"
	user := uuid.New().String()
	app := uuid.New().String()

	resp, err := Create(context.Background(), secret, user, app)
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	resp, err = GetUserSecret(context.Background(), user, app)
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, resp, secret)
	}

	resp, err = DeleteUserSecret(context.Background(), user, app)
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}
}

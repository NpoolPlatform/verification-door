package redis

import (
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

func TestRedis(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	client := Client()
	assert.Nil(t, client)

	userID := uuid.New().String()

	err := InsertKeyInfo(userID, "test", "this is a test", 0)
	assert.Nil(t, err)
}

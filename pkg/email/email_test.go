package email

import (
	"log"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init"
	"github.com/aws/aws-sdk-go/aws"
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

func TestEmail(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := SendEmailByAWS("test", "test", "crazyzplzpl@gmail.com", "crazyzplzpl@gmail.com", aws.String("crazyzplzpl@qq.com"))
	assert.Nil(t, err)
}

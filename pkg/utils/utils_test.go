package utils

import (
	"log"
	"os"
	"strconv"
	"testing"

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

func TestUtils(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	phone := "+12312312313"
	match := ValidPhoneNumber(phone)
	assert.Equal(t, true, match)

	phone = "12312313123"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "1312sadasdas"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "asdadadadadas"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "+++++++++"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "+"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "12231"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "+12545456544565445624"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)

	phone = "+12"
	match = ValidPhoneNumber(phone)
	assert.Equal(t, false, match)
}

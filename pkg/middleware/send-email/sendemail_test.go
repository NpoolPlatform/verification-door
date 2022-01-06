package sendemail

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
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

func TestSendEmailMiddleware(t *testing.T) { // nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	email := "crazyzplzpl@gmail.com"
	resp, err := SendVerifyCode(context.Background(), &npool.SendEmailRequest{
		Email:    email,
		Lang:     En,
		Username: "Crazyzpl",
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}

	resp1, err := SendUserSiteContactEmail(context.Background(), &npool.SendUserSiteContactEmailRequest{
		From:    email,
		To:      "crazyzplzpl@163.com",
		Subject: "test",
		Text:    "dsadsajlkdjaskdlasjdkalsjdasldjasildjasdjasdadsalij",
		AppID:   "test",
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp1)
	}
}

package googlerecaptcha

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/verification-door/message/npool"
	testinit "github.com/NpoolPlatform/verification-door/pkg/test-init" //nolint
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

func TestGoogleRecaptchaMiddle(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	response := "03AGdBq24Q-HKk13drAyHSYS9UYdnwxr2RhhA4df7suh34zyk9A0gv-tvYkjI0CSXEH_EjFakECJzIai08NmygNSIM_P7UC-iBrb3PBw03GKrhvh5If_a0cRaakKFrJHc0Ce7ibUFqBQMUV5XQz5_JeNibCvf1sSZHc7QIG2goPLXDORwmwJxREC4gCy_8YIa4-YxwzFjbM1MHRl8SS4kLaHKpRTm53P8uPfI4xQ4yaVf81m8PigKkGcQ9srgQxCxzfe6_eEbC4e0187HBJ-ruXLuVp0GsD4KTfFfR500xt6yj0hy1CWJfsBbBbv-ezLR32KZQ1GK3qUawjg3gwN49XDKbnT46j_ZTJYubLjsypHrIT_4_-9ul2l0EcnLUG6Xli4gJ3NoWDHVqDi3Cs6VpQPrc_TnaCdaVZzXi1rKMHKuow_r4d2ZPup6J0FtauGyuFXfMi9vvfkIzIV9UNC4JVr0iooyq50pThA"
	_, err := VerifyGoogleRecaptcha(&npool.VerifyGoogleRecaptchaRequest{
		Response: response,
	})
	assert.NotNil(t, err)
}

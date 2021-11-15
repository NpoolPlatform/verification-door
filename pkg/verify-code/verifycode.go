package verifycode

import (
	"math/rand"
	"time"

	myRedis "github.com/NpoolPlatform/verification-door/pkg/redis"
	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

const (
	VerifyCodeDuration      = 5 * time.Minute
	VerificationCodeKeyword = "verify-code"
)

func GenerateVerifyCode(length int) string {
	number := []byte("0123456789")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, number[r.Intn(len(number))])
	}
	return string(result)
}

func SaveVerifyCode(param, code string, sendTime int64) error {
	info, err := myRedis.QueryVerifyCodeKeyInfo(param, VerificationCodeKeyword)
	if err != nil && err != redis.Nil {
		return xerrors.Errorf("fail to get user verify code key: %v", err)
	}

	if err == nil {
		if (sendTime - info.SendTime) < 60*int64(time.Second) {
			return xerrors.Errorf("please wait for 60 seconds.")
		}
	}

	userCode := myRedis.VerifyUserCode{
		Code:     code,
		SendTime: sendTime,
	}

	err = myRedis.InsertKeyInfo(param, VerificationCodeKeyword, userCode, VerifyCodeDuration)
	if err != nil {
		return xerrors.Errorf("insert verify code error: %v", err)
	}

	return nil
}

func VerifyCode(param, codeInput string) error {
	info, err := myRedis.QueryVerifyCodeKeyInfo(param, VerificationCodeKeyword)
	if err != nil {
		return err
	}
	if codeInput != info.Code {
		return xerrors.Errorf("input code is wrong!")
	}
	err = myRedis.DelKey(param, VerificationCodeKeyword)
	if err != nil {
		return err
	}
	return nil
}

package verifycode

import (
	"context"
	"math/rand"
	"time"

	"github.com/NpoolPlatform/verification-door/message/npool"
	myRedis "github.com/NpoolPlatform/verification-door/pkg/redis"
	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"
)

const (
	VerifyCodeDuration      = 10 * time.Minute
	VerificationCodeKeyword = "verify-code"
	WaitTime                = 60
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

func SaveVerifyCode(ctx context.Context, param, code string, sendTime int64) error {
	info, err := myRedis.QueryVerifyCodeKeyInfo(ctx, param, VerificationCodeKeyword)
	if err != nil && err != redis.Nil {
		return xerrors.Errorf("fail to get user verify code key: %v", err)
	}

	if err == nil {
		if (sendTime - info.SendTime) < WaitTime {
			return xerrors.Errorf("please wait for 60 seconds.")
		}
	}

	userCode := myRedis.VerifyUserCode{
		Code:     code,
		SendTime: sendTime,
	}

	err = myRedis.InsertKeyInfo(ctx, param, VerificationCodeKeyword, userCode, VerifyCodeDuration)
	if err != nil {
		return xerrors.Errorf("insert verify code error: %v", err)
	}

	return nil
}

func VerifyCode(ctx context.Context, in *npool.VerifyCodeRequest) (*npool.VerifyCodeResponse, error) {
	info, err := myRedis.QueryVerifyCodeKeyInfo(ctx, in.Param, VerificationCodeKeyword)
	if err == redis.Nil {
		return nil, xerrors.Errorf("input code is wrong or expired")
	}
	if err != nil {
		return nil, err
	}

	if in.Code != info.Code {
		return nil, xerrors.Errorf("input code is wrong!")
	}
	err = myRedis.DelKey(in.Param, VerificationCodeKeyword)
	if err != nil {
		return nil, err
	}
	return &npool.VerifyCodeResponse{
		Info: "verify code successfully",
	}, nil
}

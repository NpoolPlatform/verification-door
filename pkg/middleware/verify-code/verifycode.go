package verifycode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/verification"
	myGrpc "github.com/NpoolPlatform/verification-door/pkg/grpc"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

func VerifyCodeWithUserID(ctx context.Context, in *npool.VerifyCodeWithUserIDRequest) (*npool.VerifyCodeWithUserIDResponse, error) { // nolint
	if in.GetParam() == "" {
		return nil, xerrors.Errorf("please input your email address or phone number")
	}
	resp, err := myGrpc.QueryUserInfo(ctx, in.GetUserID())
	if err != nil {
		return nil, err
	}

	if resp.EmailAddress != in.GetParam() && resp.PhoneNumber != in.GetParam() {
		return nil, xerrors.Errorf("phone or email is not binded to this user")
	}

	_, err = verifycode.VerifyCode(ctx, &npool.VerifyCodeRequest{
		Param: in.GetParam(),
		Code:  in.GetCode(),
	})
	if err != nil {
		return nil, err
	}

	return &npool.VerifyCodeWithUserIDResponse{
		Info: "successfully verify code",
	}, nil
}

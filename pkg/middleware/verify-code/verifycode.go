package verifycode

import (
	"github.com/NpoolPlatform/verification-door/message/npool"
	myGrpc "github.com/NpoolPlatform/verification-door/pkg/grpc"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"golang.org/x/xerrors"
)

func VerifyCodeWithUserID(in *npool.VerifyCodeWithUserIDRequest) (*npool.VerifyCodeWithUserIDResponse, error) { // nolint
	resp, err := myGrpc.QueryUserInfo(in.UserID)
	if err != nil {
		return nil, err
	}

	if resp.EmailAddress != in.Param && resp.PhoneNumber != in.Param {
		return nil, xerrors.Errorf("phone or email is not binded to this user")
	}

	_, err = verifycode.VerifyCode(&npool.VerifyCodeRequest{
		Param: in.Param,
		Code:  in.Code,
	})
	if err != nil {
		return nil, err
	}

	return &npool.VerifyCodeWithUserIDResponse{
		Info: "successfully verify code",
	}, nil
}

package googleauth

import (
	"context"

	"github.com/NpoolPlatform/verification-door/message/npool"
	usersecret "github.com/NpoolPlatform/verification-door/pkg/crud/user-secret"
	"github.com/NpoolPlatform/verification-door/pkg/google"
	"github.com/NpoolPlatform/verification-door/pkg/grpc"
	"golang.org/x/xerrors"
)

func GetQRcodeURL(ctx context.Context, in *npool.GetQRcodeURLRequest) (*npool.GetQRcodeURLResponse, error) {
	secret, err := google.GetSecret()
	if err != nil {
		return nil, xerrors.Errorf("fail to generate secret: %v", err)
	}

	qrcodeURL := google.GetQrcodeURL(in.Username, secret)
	_, err = usersecret.Create(ctx, secret, in.UserID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("fail to add user secret: %v", err)
	}

	return &npool.GetQRcodeURLResponse{
		Info: qrcodeURL,
	}, nil
}

func VerifyGoogleAuth(ctx context.Context, in *npool.VerifyGoogleAuthRequest) (*npool.VerifyGoogleAuthResponse, error) {
	secret, err := usersecret.GetUserSecret(ctx, in.UserID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user secret")
	}

	if result, err := google.VerifyCode(secret, in.Code); err != nil || !result {
		return nil, xerrors.Errorf("fail to verify code: %v", err)
	}

	query, err := grpc.QueryAppUser(in.AppID, in.UserID)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user app info: %v", err)
	}

	if !query.Info.GAVerify {
		err := grpc.UpdateUserGaStatus(in.UserID, in.AppID)
		if err != nil {
			return nil, xerrors.Errorf("fail to update user ga status: %v", err)
		}
	}
	return &npool.VerifyGoogleAuthResponse{
		Info: "google authentication verify successfully",
	}, nil
}

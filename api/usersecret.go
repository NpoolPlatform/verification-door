package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	usersecret "github.com/NpoolPlatform/verification-door/pkg/crud/user-secret"
	googleauth "github.com/NpoolPlatform/verification-door/pkg/middleware/google-auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetQRcodeURL(ctx context.Context, in *npool.GetQRcodeURLRequest) (*npool.GetQRcodeURLResponse, error) {
	resp, err := googleauth.GetQRcodeURL(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to get qrcode url: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) VerifyGoogleAuth(ctx context.Context, in *npool.VerifyGoogleAuthRequest) (*npool.VerifyGoogleAuthResponse, error) {
	resp, err := googleauth.VerifyGoogleAuth(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to verify google auth: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) DeleteUserGoogleAuth(ctx context.Context, in *npool.DeleteUserGoogleAuthRequest) (*npool.DeleteUserGoogleAuthResponse, error) {
	resp, err := usersecret.DeleteUserSecret(ctx, in.UserID)
	if err != nil {
		logger.Sugar().Errorf("fail to delete google auth: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &npool.DeleteUserGoogleAuthResponse{
		Info: resp,
	}, nil
}

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	middleware "github.com/NpoolPlatform/verification-door/pkg/middleware/verify-code"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyCode(ctx context.Context, in *npool.VerifyCodeRequest) (*npool.VerifyCodeResponse, error) {
	resp, err := verifycode.VerifyCode(in)
	if err != nil {
		logger.Sugar().Errorf("fail to verify code: %v", err)
		return &npool.VerifyCodeResponse{}, status.Errorf(codes.FailedPrecondition, "Internal server error: %v", err)
	}
	return resp, nil
}

func (s *Server) VerifyCodeWithUserID(ctx context.Context, in *npool.VerifyCodeWithUserIDRequest) (*npool.VerifyCodeWithUserIDResponse, error) {
	resp, err := middleware.VerifyCodeWithUserID(in)
	if err != nil {
		logger.Sugar().Errorf("fail to verify code: %v", err)
		return &npool.VerifyCodeWithUserIDResponse{}, status.Errorf(codes.FailedPrecondition, "Internal server error: %v", err)
	}
	return resp, nil
}

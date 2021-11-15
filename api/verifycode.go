package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyCode(ctx context.Context, in *npool.VerifyCodeRequest) (*npool.VerifyCodeResponse, error) {
	err := verifycode.VerifyCode(in.Param, in.Code)
	if err != nil {
		logger.Sugar().Errorf("fail to verify code: %v", err)
		return &npool.VerifyCodeResponse{}, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &npool.VerifyCodeResponse{
		Info: "verify code successfully",
	}, nil
}

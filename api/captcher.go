package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/NpoolPlatform/verification-door/pkg/captcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCaptcherImg(ctx context.Context, in *npool.GetCaptcherImgRequest) (*npool.GetCaptcherImgResponse, error) {
	resp, err := captcher.GetCaptcherImg()
	if err != nil {
		logger.Sugar().Errorf("fail to get captcher img: %v", err)
		return &npool.GetCaptcherImgResponse{}, status.Errorf(codes.FailedPrecondition, "failed: %v", err)
	}
	return resp, nil
}

func (s *Server) VerifyCaptcher(ctx context.Context, in *npool.VerifyCaptcherRequest) (*npool.VerifyCaptcherResponse, error) {
	err := captcher.VerifyCaptcher(in)
	if err != nil {
		logger.Sugar().Errorf("fail to verify captcher: %v", err)
		return &npool.VerifyCaptcherResponse{}, status.Errorf(codes.FailedPrecondition, "failed: %v", err)
	}
	return &npool.VerifyCaptcherResponse{
		Info: "verify successfully",
	}, nil
}

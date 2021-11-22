package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	googlerecaptcha "github.com/NpoolPlatform/verification-door/pkg/middleware/google-recaptcha"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyGoogleRecaptcha(ctx context.Context, in *npool.VerifyGoogleRecaptchaRequest) (*npool.VerifyGoogleRecaptchaResponse, error) {
	resp, err := googlerecaptcha.VerifyGoogleRecaptcha(in)
	if err != nil {
		logger.Sugar().Errorf("fail to verify google recaptcha: %v", err)
		return resp, status.Errorf(codes.Internal, "internale server error: %v", err)
	}
	return resp, err
}

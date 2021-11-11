package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	sendemail "github.com/NpoolPlatform/verification-door/pkg/middleware/send-email"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendEmail(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	resp, err := sendemail.SendEmail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to send email: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}

	return resp, nil
}

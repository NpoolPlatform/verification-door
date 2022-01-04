package api

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/verification-door/message/npool"
	sendemail "github.com/NpoolPlatform/verification-door/pkg/middleware/send-email"
	sendsms "github.com/NpoolPlatform/verification-door/pkg/middleware/send-sms"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendEmail(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	if in.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "email address cannot be null")
	}
	resp, err := sendemail.SendVerifyCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to send email: %v", err)
		return nil, status.Errorf(codes.FailedPrecondition, "Internal server error: %v", err)
	}

	return resp, nil
}

func (s *Server) SendSms(ctx context.Context, in *npool.SendSmsRequest) (*npool.SendSmsResponse, error) {
	resp, err := sendsms.SendVerifyCode(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to send sms message: %v", err)
		return nil, status.Errorf(codes.FailedPrecondition, "fail to send sms message: %v", err)
	}
	return resp, nil
}

func (s *Server) SendUserSiteContactEmail(ctx context.Context, in *npool.SendUserSiteContactEmailRequest) (*npool.SendUserSiteContactEmailResponse, error) {
	if in.From == "" || in.To == "" {
		logger.Sugar().Error("fail to send user site contact email: Sender email address and receiver email address cannot be null.")
		return nil, status.Error(codes.InvalidArgument, "sender email address and receiver email address cannot be null.")
	}

	if in.Text == "" {
		logger.Sugar().Error("fail to send user site contact email: email text cannot be null")
		return nil, status.Error(codes.InvalidArgument, "email text cannot be null, please enter your contact message")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := sendemail.SendUserSiteContactEmail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail to send user site contact email: %v", err)
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return resp, nil
}

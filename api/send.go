package api

import (
	"context"
	"net/mail"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/verification"
	sendemail "github.com/NpoolPlatform/verification-door/pkg/middleware/send-email"
	sendsms "github.com/NpoolPlatform/verification-door/pkg/middleware/send-sms"
	"github.com/NpoolPlatform/verification-door/pkg/utils"
	verifycode "github.com/NpoolPlatform/verification-door/pkg/verify-code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SendEmail(ctx context.Context, in *npool.SendEmailRequest) (*npool.SendEmailResponse, error) {
	if in.GetEmail() == "" {
		logger.Sugar().Error("SendEmail error: Email address cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "email address cannot be empty")
	}

	if _, err := mail.ParseAddress(in.GetEmail()); err != nil {
		logger.Sugar().Errorf("SendEmail error, invalid email address: %v", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid email address")
	}

	resp, err := sendemail.SendVerifyCode(ctx, in)
	if err != nil {
		if err.Error() == verifycode.WaitTimeError {
			logger.Sugar().Errorf("SendEmail error: %v", err)
			return nil, status.Errorf(codes.AlreadyExists, "%v", err)
		}

		logger.Sugar().Errorf("SendEmail error: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return resp, nil
}

func (s *Server) SendSms(ctx context.Context, in *npool.SendSmsRequest) (*npool.SendSmsResponse, error) {
	if in.GetPhone() == "" {
		logger.Sugar().Error("SendSms error: phone number can not be empty")
		return nil, status.Error(codes.InvalidArgument, "phone number can not be empty")
	}

	if match := utils.ValidPhoneNumber(in.GetPhone()); !match {
		logger.Sugar().Error("SendSms error: phone number is not invalid")
		return nil, status.Error(codes.InvalidArgument, "phone number is not invalid")
	}

	resp, err := sendsms.SendVerifyCode(ctx, in)
	if err != nil {
		if err.Error() == verifycode.WaitTimeError {
			logger.Sugar().Errorf("SendSms error: %v", err)
			return nil, status.Errorf(codes.AlreadyExists, "%v", err)
		}

		logger.Sugar().Errorf("SendSms error, internal server error: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return resp, nil
}

func (s *Server) SendUserSiteContactEmail(ctx context.Context, in *npool.SendUserSiteContactEmailRequest) (*npool.SendUserSiteContactEmailResponse, error) {
	if in.GetFrom() == "" {
		logger.Sugar().Error("SendUserSiteContactEmail error: 'From' email address can not be empty")
		return nil, status.Error(codes.InvalidArgument, "'From' email address can not be empty")
	}

	if in.GetTo() == "" {
		logger.Sugar().Error("SendUserSiteContactEmail error: 'To' email address can not be empty")
		return nil, status.Error(codes.InvalidArgument, "'To' email address can not be empty")
	}

	if _, err := mail.ParseAddress(in.GetFrom()); err != nil {
		logger.Sugar().Error("SendUserSiteContactEmail error: 'From' email address is not valid")
		return nil, status.Error(codes.InvalidArgument, "'From' email address is not valid")
	}

	if _, err := mail.ParseAddress(in.GetTo()); err != nil {
		logger.Sugar().Error("SendUserSiteContactEmail error: 'To' email address is not valid")
		return nil, status.Error(codes.InvalidArgument, "'To' email address is not valid")
	}

	if in.GetText() == "" {
		logger.Sugar().Error("SendUserSiteContactEmail error: email text cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "email text can not be empty")
	}

	if in.GetSubject() == "" {
		logger.Sugar().Error("SendUserSiteContactEmail error: email subject can not be empty")
		return nil, status.Error(codes.InvalidArgument, "email subject can not be empty")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := sendemail.SendUserSiteContactEmail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("SendUserSiteContactEmail error, internal server error: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return resp, nil
}

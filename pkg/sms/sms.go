package sms

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	Region    = "aws_region"
	AccessKey = "aws_access_key"
	SecretKey = "aws_secret_key"
)

func SendSms(msg, phone string) error {
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	region := config.GetStringValueWithNameSpace(myServiceName, Region)
	accessKey := config.GetStringValueWithNameSpace(myServiceName, AccessKey)
	secretKey := config.GetStringValueWithNameSpace(myServiceName, SecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return err
	}

	pin := &sns.PublishInput{}
	pin.SetMessage(msg)
	pin.SetPhoneNumber(phone)

	svc := sns.New(sess)
	_, err = svc.Publish(pin)
	if err != nil {
		return err
	}

	return nil
}

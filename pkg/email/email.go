package email

import (
	"net/mail"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"golang.org/x/xerrors"
)

const (
	Region    = "aws_region"
	AccessKey = "aws_access_key"
	SecretKey = "aws_secret_key"
	CharSet   = "UTF-8"
)

func SendEmailByAWS(subtitle, content, from, to string, replyTo ...*string) error {
	_, err := mail.ParseAddress(to)
	if err != nil {
		return xerrors.Errorf("invalid email address: %v", err)
	}

	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	region := config.GetStringValueWithNameSpace(myServiceName, Region)
	accessKey := config.GetStringValueWithNameSpace(myServiceName, AccessKey)
	secretKey := config.GetStringValueWithNameSpace(myServiceName, SecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return xerrors.Errorf("new aws session error: %v", err)
	}

	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(content),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subtitle),
			},
		},
		Source: aws.String(from),
	}

	if len(replyTo) != 0 {
		input.ReplyToAddresses = replyTo
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return xerrors.Errorf("%v, %v", ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return xerrors.Errorf("%v, %v", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return xerrors.Errorf("%v, %v", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				return aerr
			}
		} else {
			return err
		}
	}

	return nil
}

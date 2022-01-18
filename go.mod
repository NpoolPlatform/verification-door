module github.com/NpoolPlatform/verification-door

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/NpoolPlatform/application-management v0.0.0-20220118102249-a68703db2408
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211222114515-4928e6cf3f1f
	github.com/NpoolPlatform/message v0.0.0-20220118090327-926885a280ec
	github.com/NpoolPlatform/user-management v0.0.0-20220118105710-fb22a0c56085
	github.com/aws/aws-sdk-go v1.42.19
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2
	github.com/mojocn/base64Captcha v1.3.5
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.43.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0

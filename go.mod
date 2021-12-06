module github.com/NpoolPlatform/verification-door

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/NpoolPlatform/application-management v0.0.0-20211206033605-504c9e2b83ec
	github.com/NpoolPlatform/go-service-framework v0.0.0-20211117074545-bc1340849b08
	github.com/aws/aws-sdk-go v1.42.19
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mailgun/mailgun-go/v3 v3.6.4
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/rogpeppe/go-internal v1.8.1-0.20211023094830-115ce09fd6b4 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/sys v0.0.0-20211110154304-99a53858aa08 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/genproto v0.0.0-20211117155847-120650a500bb
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.41.0

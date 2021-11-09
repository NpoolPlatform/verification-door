package api

import (
	"context"

	"github.com/NpoolPlatform/verification-door/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false
type Server struct {
	npool.UnimplementedVerificationDoorServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterVerificationDoorServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterVerificationDoorHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

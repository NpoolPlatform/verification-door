package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
)

func newApplicationGrpcConn() (*grpc.ClientConn, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func QueryAppUser(appID, userID string) (*pbApplication.GetUserFromApplicationResponse, error) {
	conn, err := newApplicationGrpcConn()
	if err != nil {
		return &pbApplication.GetUserFromApplicationResponse{}, err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	resp, err := client.GetUserFromApplication(context.Background(), &pbApplication.GetUserFromApplicationRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return &pbApplication.GetUserFromApplicationResponse{}, err
	}

	return resp, nil
}

func UpdateUserGaStatus(userID, appID string) error {
	conn, err := newApplicationGrpcConn()
	if err != nil {
		return err
	}

	client := pbApplication.NewApplicationManagementClient(conn)
	_, err = client.UpdateUserGAStatus(context.Background(), &pbApplication.UpdateUserGAStatusRequest{
		UserID: userID,
		AppID:  appID,
		Status: true,
	})
	if err != nil {
		return err
	}

	return nil
}

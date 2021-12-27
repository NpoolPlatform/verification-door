package grpc

import (
	"context"
	"time"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbuser "github.com/NpoolPlatform/user-management/message/npool"
	userconst "github.com/NpoolPlatform/user-management/pkg/message/const"
)

const (
	grpcTimeout = 5 * time.Second
)

func QueryAppUser(ctx context.Context, appID, userID string) (*pbApplication.GetUserFromApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	resp, err := client.GetUserFromApplication(ctx, &pbApplication.GetUserFromApplicationRequest{
		AppID:  appID,
		UserID: userID,
	})
	if err != nil {
		return &pbApplication.GetUserFromApplicationResponse{}, err
	}

	return resp, nil
}

func UpdateUserGaStatus(ctx context.Context, userID, appID string) error {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pbApplication.NewApplicationManagementClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	_, err = client.UpdateUserGAStatus(ctx, &pbApplication.UpdateUserGAStatusRequest{
		UserID: userID,
		AppID:  appID,
		Status: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func QueryUserInfo(ctx context.Context, userID string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pbuser.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	resp, err := client.GetUser(ctx, &pbuser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Info, nil
}

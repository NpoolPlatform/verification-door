package grpc

import (
	"context"

	pbApplication "github.com/NpoolPlatform/application-management/message/npool"
	applicationconst "github.com/NpoolPlatform/application-management/pkg/message/const"
	mygrpc "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	pbuser "github.com/NpoolPlatform/user-management/message/npool"
	userconst "github.com/NpoolPlatform/user-management/pkg/message/const"
)

func QueryAppUser(appID, userID string) (*pbApplication.GetUserFromApplicationResponse, error) {
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
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
	conn, err := mygrpc.GetGRPCConn(applicationconst.ServiceName, mygrpc.GRPCTAG)
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

func QueryUserInfo(userID string) (*pbuser.UserBasicInfo, error) {
	conn, err := mygrpc.GetGRPCConn(userconst.ServiceName, mygrpc.GRPCTAG)
	if err != nil {
		return nil, err
	}

	client := pbuser.NewUserClient(conn)
	resp, err := client.GetUser(context.Background(), &pbuser.GetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Info, nil
}

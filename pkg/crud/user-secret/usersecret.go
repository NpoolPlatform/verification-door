package usersecret

import (
	"context"
	"time"

	"github.com/NpoolPlatform/verification-door/pkg/db"
	"github.com/NpoolPlatform/verification-door/pkg/db/ent/usersecret"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func Create(ctx context.Context, secret, userID, appID string) (string, error) {
	user, err := uuid.Parse(userID)
	if err != nil {
		return "", xerrors.Errorf("invalid user id: %v", err)
	}

	app, err := uuid.Parse(appID)
	if err != nil {
		return "", xerrors.Errorf("invalid app id: %v", err)
	}

	info, err := db.Client().
		UserSecret.
		Query().
		Where(
			usersecret.And(
				usersecret.UserID(user),
				usersecret.AppID(app),
				usersecret.DeleteAt(0),
			),
		).All(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to get user secret record: %v", err)
	}

	if len(info) != 0 {
		return "", xerrors.Errorf("user already has a google auth secret")
	}

	_, err = db.Client().
		UserSecret.
		Create().
		SetUserID(user).
		SetAppID(app).
		SetSecret(secret).
		Save(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to create user secret to mysql: %v", err)
	}
	return "successfully", nil
}

func GetUserSecret(ctx context.Context, userID, appID string) (string, error) {
	user, err := uuid.Parse(userID)
	if err != nil {
		return "", xerrors.Errorf("invalid user id: %v", err)
	}

	app, err := uuid.Parse(appID)
	if err != nil {
		return "", xerrors.Errorf("invalid app id: %v", err)
	}

	info, err := db.Client().
		UserSecret.
		Query().
		Where(
			usersecret.And(
				usersecret.UserID(user),
				usersecret.AppID(app),
				usersecret.DeleteAt(0),
			),
		).Only(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to get user secret: %v", err)
	}

	return info.Secret, nil
}

func DeleteUserSecret(ctx context.Context, userID, appID string) (string, error) {
	user, err := uuid.Parse(userID)
	if err != nil {
		return "", xerrors.Errorf("invalid user id: %v", err)
	}

	app, err := uuid.Parse(appID)
	if err != nil {
		return "", xerrors.Errorf("invalid app id: %v", err)
	}

	_, err = db.Client().
		UserSecret.
		Update().
		Where(
			usersecret.And(
				usersecret.UserID(user),
				usersecret.AppID(app),
				usersecret.DeleteAt(0),
			),
		).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return "", xerrors.Errorf("fail to delete user secret: %v", err)
	}

	return "delete successfully", nil
}

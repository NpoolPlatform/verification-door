package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	myRedis "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/go-redis/redis/v8"
)

type VerifyUserCode struct {
	Code     string `json:"code"`
	SendTime int64  `json:"send_time"`
}

func Client() *redis.Client {
	return myRedis.Client{}.Client
}

func InsertKeyInfo(userID, keyWord string, info interface{}, ttl time.Duration) error {
	b, err := json.Marshal(info)
	if err != nil {
		return err
	}
	err = Client().Set(context.Background(), fmt.Sprintf("%v::%v", keyWord, userID), string(b), ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func QueryVerifyCodeKeyInfo(userID, keyWord string) (*VerifyUserCode, error) {
	val, err := Client().Get(context.Background(), fmt.Sprintf("%v::%v", keyWord, userID)).Result()
	if err != nil {
		return nil, err
	}
	info := &VerifyUserCode{}
	err = json.Unmarshal([]byte(val), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

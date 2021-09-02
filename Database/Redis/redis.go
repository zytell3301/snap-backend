package Redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Configs *viper.Viper
var Connection *redis.Client

func InitiateRedis() {
	connection := redis.NewClient(&redis.Options{
		Addr:      Configs.GetString("host") + ":" + Configs.GetString("port"),
		OnConnect: onConnect,
	})

	Connection = connection
}

func onConnect(_ context.Context, cn *redis.Conn) error {
	status := cn.Ping(context.Background())
	result, _ := status.Result()

	switch result != "PONG" {
	case true:
		panic("Failed to ping redis server")
	}
	return nil
}

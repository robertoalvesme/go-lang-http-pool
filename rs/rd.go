package rs

import (
	"github.com/hoisie/redis"
	"gitlab.com/begrowth/connect/service-mail-verify/src/infrastructure/config"
	"log"
)

func GetRedisClient() redis.Client {
	client := redis.Client{
		Addr:     config.REDIS_HOST + ":" + config.REDIS_PORT,
		Password: config.REDIS_PASS,
		Db:       0,
	}

	log.Println("Reddis Client Connected")
	return client
}

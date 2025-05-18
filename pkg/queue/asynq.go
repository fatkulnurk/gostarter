package queue

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"magicauth/config"
)

func NewAsynqClient(cfg *config.Queue, redis *redis.Client) (*asynq.Client, error) {
	client := asynq.NewClientFromRedisClient(redis)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	err := client.Ping()
	if err != nil {
		return nil, err
	}

	return client, nil
}

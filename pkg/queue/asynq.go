package queue

import (
	"context"
	"encoding/json"

	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/pkg/logging"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

func NewAsynqClient(cfg *config.Queue, redis *redis.Client) (*asynq.Client, error) {
	client := asynq.NewClientFromRedisClient(redis)

	err := client.Ping()
	if err != nil {
		logging.Fatalf("failed to ping redis: %v", err)
		return nil, err
	}

	return client, nil
}

type AsynqQueue struct {
	client *asynq.Client
}

func NewAsynqQueue(client *asynq.Client) IQueue {
	return &AsynqQueue{client: client}
}

func (q *AsynqQueue) Enqueue(ctx context.Context, taskName string, payload any, opts ...Option) (*OutputEnqueue, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(taskName, data)
	aOpts := toAsynqOptions(opts...)
	tInfo, err := q.client.EnqueueContext(ctx, task, aOpts...)
	if err != nil {
		return nil, err
	}
	return &OutputEnqueue{TaskID: tInfo.ID, Payload: data, Options: opts}, nil
}

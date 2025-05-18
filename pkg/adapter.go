package pkg

import (
	"database/sql"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

//Driver adalah input side — seperti handler HTTP, CLI, scheduler, dsb.
//Adapter adalah penghubung antara domain/usecase dengan infrastruktur, seperti database, HTTP client, file storage, dll.

type Adapter struct {
	DB    *sql.DB
	Redis *redis.Client
	Queue *asynq.Client
}

func NewAdapter(db *sql.DB, redis *redis.Client, queue *asynq.Client) *Adapter {
	return &Adapter{
		DB:    db,
		Redis: redis,
		Queue: queue,
	}
}

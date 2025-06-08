package pkg

import (
	"database/sql"

	"github.com/fatkulnurk/gostarter/pkg/cache"
	"github.com/fatkulnurk/gostarter/pkg/mailer"
	"github.com/fatkulnurk/gostarter/pkg/queue"
	"github.com/fatkulnurk/gostarter/pkg/storage"
	"github.com/redis/go-redis/v9"
)

//Driver adalah input side — seperti handler HTTP, CLI, scheduler, dsb.
//Adapter adalah penghubung antara domain/usecase dengan infrastruktur, seperti database, HTTP client, file storage, dll.

type Database struct {
	DB    *sql.DB
	Redis *redis.Client
}

type Adapter struct {
	DB      *Database
	Cache   *cache.Cache
	Mailer  *mailer.Mailer
	Queue   *queue.Queue
	Storage *storage.Storage
}

func NewAdapter(db *Database, cache *cache.Cache, mailer *mailer.Mailer, queue *queue.Queue, storage *storage.Storage) *Adapter {
	return &Adapter{
		DB:      db,
		Cache:   cache,
		Mailer:  mailer,
		Queue:   queue,
		Storage: storage,
	}
}

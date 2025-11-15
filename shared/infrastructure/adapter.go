package infrastructure

import (
	"database/sql"

	"github.com/fatkulnurk/gostarter/pkg/cache"
	"github.com/fatkulnurk/gostarter/pkg/mailer"
	"github.com/fatkulnurk/gostarter/pkg/queue"
	"github.com/fatkulnurk/gostarter/pkg/storage"
	"github.com/redis/go-redis/v9"
)

// Clean Architecture Infrastructure Layer:
// - Driver: Input side components like HTTP handlers, CLI commands, schedulers, etc.
// - Adapter: Interface between domain/business logic and infrastructure like databases, HTTP clients, file storage, etc.

// DatabaseConnection holds all database connections used by the application
// It provides access to SQL and Redis databases for the application
type DatabaseConnection struct {
	Sql   *sql.DB
	Redis *redis.Client
}

// Adapter provides access to all infrastructure components needed by the application
// It implements the adapter pattern from clean architecture to abstract infrastructure details from business logic
// This allows domain logic to remain independent of infrastructure concerns
type Adapter struct {
	DB      *DatabaseConnection
	Cache   *cache.Cache
	Mailer  *mailer.Mailer
	Queue   *queue.Queue
	Storage *storage.Storage
}

// NewAdapter creates a new Adapter instance with all required infrastructure dependencies
// This function centralizes the creation of the adapter to ensure all required dependencies are provided
func NewAdapter(db *DatabaseConnection, cache *cache.Cache, mailer *mailer.Mailer, queue *queue.Queue, storage *storage.Storage) *Adapter {
	return &Adapter{
		DB:      db,
		Cache:   cache,
		Mailer:  mailer,
		Queue:   queue,
		Storage: storage,
	}
}

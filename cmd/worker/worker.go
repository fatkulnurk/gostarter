package worker

import (
	"fmt"
	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/internal/example"
	"github.com/fatkulnurk/gostarter/pkg/db"
	"github.com/fatkulnurk/gostarter/pkg/module"
	pkgqueue "github.com/fatkulnurk/gostarter/pkg/queue"
	"github.com/fatkulnurk/gostarter/shared/infrastructure"
	"github.com/hibiken/asynq"
	"log"
)

func Serve(cfg *config.Config) {
	// adapter, only register what you need
	adapter := func(cfg *config.Config) *infrastructure.Adapter {
		mysql, err := db.NewMySQL(cfg.Database)
		if err != nil {
			panic(err)
		}

		redis, err := db.NewRedis(cfg.Redis)
		if err != nil {
			panic(err)
		}

		asynqClient, err := pkgqueue.NewAsynqClient(cfg.Queue, redis)
		if err != nil {
			panic(err)
		}

		queue := pkgqueue.NewAsynqQueue(asynqClient)
		return &infrastructure.Adapter{
			DB: &infrastructure.DatabaseConnection{
				Redis: redis,
				Sql:   mysql,
			},
			Queue: &queue,
		}
	}(cfg)

	// delivery, only register what you need
	delivery := func(cfg *config.Config) *infrastructure.Delivery {
		mux := asynq.NewServeMux()
		return &infrastructure.Delivery{
			HTTP: nil,
			Task: mux,
		}
	}(cfg)

	// Register modules
	func() {
		var modules []module.IModule
		modules = append(modules, example.New(adapter, delivery))

		fmt.Printf("-------Register mdl------\n")
		for idx, mdl := range modules {
			fmt.Printf("number: %d\n", idx+1)
			fmt.Printf("Registering mdl: %s\n", mdl.GetInfo().Name)
			fmt.Printf("Prefix: %s\n", mdl.GetInfo().Prefix)
			mdl.RegisterTask()
			fmt.Printf("-------------------------\n")
		}
	}()

	server := asynq.NewServerFromRedisClient(adapter.DB.Redis,
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: cfg.Queue.Concurrency,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	if err := server.Run(delivery.Task); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

package worker

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"magicauth/config"
	"magicauth/internal/example"
	"magicauth/pkg"
	"magicauth/pkg/db"
	pkgqueue "magicauth/pkg/queue"
)

func Serve(cfg *config.Config) {
	// adapter, only register what you need
	adapter := func(cfg *config.Config) *pkg.Adapter {
		mysql, err := db.NewMySQL(cfg.Database)
		if err != nil {
			panic(err)
		}

		redis, err := db.NewRedis(cfg.Redis)
		if err != nil {
			panic(err)
		}

		queue, err := pkgqueue.NewAsynqClient(cfg.Queue, redis)
		if err != nil {
			panic(err)
		}

		return &pkg.Adapter{
			DB:    mysql,
			Redis: redis,
			Queue: queue,
		}
	}(cfg)

	// delivery, only register what you need
	delivery := func(cfg *config.Config, adapter *pkg.Adapter) *pkg.Delivery {
		mux := asynq.NewServeMux()
		scheduler := asynq.NewSchedulerFromRedisClient(adapter.Redis, nil)
		return &pkg.Delivery{
			Task:     mux,
			Schedule: scheduler,
		}
	}(cfg, adapter)

	// Register modules
	func() {
		var modules []pkg.IModule
		modules = append(modules, example.New(adapter, delivery))

		fmt.Printf("-------Register module------\n")
		for idx, module := range modules {
			fmt.Printf("number: %d\n", idx+1)
			fmt.Printf("Registering module: %s\n", module.GetInfo().Name)
			fmt.Printf("Prefix: %s\n", module.GetInfo().Prefix)
			module.RegisterSchedule()
			fmt.Printf("-------------------------\n")
		}
	}()

	if err := delivery.Schedule.Run(); err != nil {
		log.Fatal(err)
	}
}

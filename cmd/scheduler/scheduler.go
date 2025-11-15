package scheduler

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
	"time"
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
				Sql:   mysql,
				Redis: redis,
			},
			Queue: &queue,
		}
	}(cfg)

	// delivery, only register what you need
	delivery := func(cfg *config.Config, adapter *infrastructure.Adapter) *infrastructure.Delivery {
		timeLocation, err := time.LoadLocation(cfg.Schedule.Timezone)
		if err != nil {
			panic(err)
		}

		mux := asynq.NewServeMux()
		scheduler := asynq.NewSchedulerFromRedisClient(adapter.DB.Redis, &asynq.SchedulerOpts{
			Location: timeLocation,
		})
		return &infrastructure.Delivery{
			Task:     mux,
			Schedule: scheduler,
		}
	}(cfg, adapter)

	// Register modules
	func() {
		var modules []module.IModule
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

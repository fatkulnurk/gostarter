package worker

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"magicauth/config"
	"magicauth/internal/magiclink"
	"magicauth/internal/twofa"
	"magicauth/pkg"
)

func Serve(cfg *config.Config) {
	// adapter, only register what you need
	adapter := func(cfg *config.Config) *pkg.Adapter {
		return &pkg.Adapter{
			DB:    nil,
			Redis: nil,
			Queue: nil,
		}
	}(cfg)

	// delivery, only register what you need
	delivery := func(cfg *config.Config) *pkg.Delivery {
		mux := asynq.NewServeMux()
		return &pkg.Delivery{
			HTTP:   nil,
			Worker: mux,
		}
	}(cfg)

	// Register modules
	func() {
		var modules []pkg.IModule
		modules = append(modules, twofa.New(adapter, delivery))
		modules = append(modules, magiclink.New(adapter, delivery))

		fmt.Printf("-------Register module------\n")
		for idx, module := range modules {
			fmt.Printf("number: %d\n", idx+1)
			fmt.Printf("Registering module: %s\n", module.GetInfo().Name)
			fmt.Printf("Prefix: %s\n", module.GetInfo().Prefix)
			module.RegisterHTTP()
			fmt.Printf("-------------------------\n")
		}
	}()

	server := asynq.NewServerFromRedisClient(adapter.Redis,
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

	if err := server.Run(delivery.Worker); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

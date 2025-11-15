package http

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatkulnurk/gostarter/pkg/config"
	"github.com/fatkulnurk/gostarter/pkg/module"
	"github.com/fatkulnurk/gostarter/shared/infrastructure"

	"github.com/fatkulnurk/gostarter/shared/middleware"

	"github.com/fatkulnurk/gostarter/internal/example"
	"github.com/fatkulnurk/gostarter/pkg/db"
	pkgqueue "github.com/fatkulnurk/gostarter/pkg/queue"
	"github.com/gofiber/fiber/v2"
	gofibermiddlewarerecover "github.com/gofiber/fiber/v2/middleware/recover"
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
	delivery := func(cfg *config.Config) *infrastructure.Delivery {
		return &infrastructure.Delivery{
			HTTP: initHttp(cfg),
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
			mdl.RegisterHTTP()
			fmt.Printf("-------------------------\n")
		}
	}()

	// Create a channel to listen for interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := delivery.HTTP.Listen(":8080"); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-c
	fmt.Println("Shutting down gracefully...")

	// Shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := delivery.HTTP.ShutdownWithContext(ctx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	}
}

func initHttp(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       cfg.DeliveryHttp.Prefork,
		CaseSensitive: cfg.DeliveryHttp.CaseSensitive,
		StrictRouting: cfg.DeliveryHttp.StrictRouting,
		ServerHeader:  cfg.DeliveryHttp.ServerHeader,
		AppName:       cfg.App.Name,
		BodyLimit:     cfg.DeliveryHttp.BodyLimit,
	})
	app.Use(gofibermiddlewarerecover.New())
	app.Use(middleware.LoggingMiddleware())
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	return app
}

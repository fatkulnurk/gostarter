package http

import (
	"context"
	"fmt"
	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/internal/example"
	"github.com/fatkulnurk/gostarter/pkg"
	"github.com/fatkulnurk/gostarter/pkg/db"
	pkgqueue "github.com/fatkulnurk/gostarter/pkg/queue"
	"github.com/gofiber/fiber/v2"
	gofibermiddlewarerecover "github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	delivery := func(cfg *config.Config) *pkg.Delivery {
		return &pkg.Delivery{
			HTTP: initHttp(cfg),
		}
	}(cfg)

	// Register modules
	func() {
		var modules []pkg.IModule
		modules = append(modules, example.New(adapter, delivery))

		fmt.Printf("-------Register module------\n")
		for idx, module := range modules {
			fmt.Printf("number: %d\n", idx+1)
			fmt.Printf("Registering module: %s\n", module.GetInfo().Name)
			fmt.Printf("Prefix: %s\n", module.GetInfo().Prefix)
			module.RegisterTask()
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
	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("[%s] %s - %s\n", c.Method(), c.Path(), c.IP())
		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("API is running")
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	return app
}

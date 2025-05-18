package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"magicauth/cmd/http"
	"magicauth/cmd/worker"
	"magicauth/config"
	"os"
	"strconv"
	"time"

	"magicauth/pkg/cache"
	"magicauth/pkg/db"

	"github.com/redis/go-redis/v9"
)

func ServeApp(svc string, cfg *config.Config) {
	switch svc {
	case "http":
		fmt.Println("Running in HTTP server mode...")
		http.Serve(cfg)
	case "worker":
		fmt.Println("Running in background worker mode...")
		worker.Serve(cfg)
	default:
		_, err := fmt.Fprintf(os.Stderr, "Error: invalid --app value: %s\n", svc)
		if err != nil {
			return
		}
		flag.Usage()
		os.Exit(1)
	}
}

func InitMySQL() (*sql.DB, error) {
	params := os.Getenv("DB_PARAMS")
	if params == "" {
		params = "charset=utf8mb4&parseTime=True&loc=Local"
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	conn, err := db.NewMySQL(db.MySQLConfig{
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		Host:            os.Getenv("DB_HOST"),
		Port:            port,
		Database:        os.Getenv("DB_DATABASE"),
		Params:          params,
		MaxOpenConns:    25,
		MaxIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func InitRedis() (*redis.Client, error) {
	return db.NewRedis(cache.RedisConfig{
		Addr:            os.Getenv("REDIS_ADDR"),
		Password:        os.Getenv("REDIS_PASSWORD"),
		DB:              0,
		PoolSize:        10,
		MinIdleConns:    5,
		ConnMaxLifetime: 5 * time.Minute,
		PoolTimeout:     4 * time.Second,
		ConnMaxIdleTime: 30 * time.Minute,
	})
}

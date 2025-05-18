package config

import (
	"time"

	"github.com/fatkulnurk/gostarter/shared/constant"
	"github.com/fatkulnurk/gostarter/shared/utils"
	"github.com/joho/godotenv"
)

func New(env string) *Config {
	// default environment is development
	envFile := ""
	if env == "" {
		env = constant.EnvironmentDevelopment
	}

	switch env {
	case constant.EnvironmentDevelopment:
		envFile = ".env"
	case constant.EnvironmentTest:
		envFile = ".env.test"
	}

	// load environment variables
	if err := godotenv.Load(envFile); err != nil {
		panic(err)
	}
	cfg := Config{
		App: &App{
			Environment: env,
			Name:        utils.GetEnv("APP_NAME", "GoStarter"),
			Version:     utils.GetEnv("APP_VERSION", "1.0.0"),
		},
		Database: &Database{
			User:            utils.GetEnv("DB_USER", "root"),
			Password:        utils.GetEnv("DB_PASSWORD", ""),
			Host:            utils.GetEnv("DB_HOST", "localhost"),
			Port:            utils.GetIntEnv("DB_PORT", 3306),
			Database:        utils.GetEnv("DB_NAME", "gostarter"),
			Params:          utils.GetEnv("DB_PARAMS", "charset=utf8mb4&parseTime=true"),
			MaxOpenConns:    utils.GetIntEnv("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:    utils.GetIntEnv("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: utils.GetDurationEnv("DB_CONN_MAX_LIFETIME", time.Hour),
			ConnMaxIdleTime: utils.GetDurationEnv("DB_CONN_MAX_IDLE_TIME", time.Minute*30),
		},
		DeliveryHttp: &DeliveryHttp{
			Prefork:       utils.GetBoolEnv("HTTP_PREFORK", false),
			CaseSensitive: utils.GetBoolEnv("HTTP_CASE_SENSITIVE", true),
			StrictRouting: utils.GetBoolEnv("HTTP_STRICT_ROUTING", false),
			BodyLimit:     utils.GetIntEnv("HTTP_BODY_LIMIT", 10*1024*1024),
			ServerHeader:  utils.GetEnv("HTTP_SERVER_HEADER", "GoStarter"),
		},
		DeliveryQueue: &DeliveryQueue{
			Concurrency: utils.GetIntEnv("QUEUE_CONCURRENCY", 10),
		},
		Redis: &Redis{
			Addr:            utils.GetEnv("REDIS_ADDR", "localhost:6379"),
			Password:        utils.GetEnv("REDIS_PASSWORD", ""),
			DB:              utils.GetIntEnv("REDIS_DB", 0),
			PoolSize:        utils.GetIntEnv("REDIS_POOL_SIZE", 10),
			MinIdleConns:    utils.GetIntEnv("REDIS_MIN_IDLE_CONNS", 5),
			ConnMaxLifetime: utils.GetDurationEnv("REDIS_CONN_MAX_LIFETIME", time.Hour),
			PoolTimeout:     utils.GetDurationEnv("REDIS_POOL_TIMEOUT", time.Second*4),
			ConnMaxIdleTime: utils.GetDurationEnv("REDIS_CONN_MAX_IDLE_TIME", time.Minute*30),
			ReadTimeout:     utils.GetDurationEnv("REDIS_READ_TIMEOUT", time.Second*3),
			WriteTimeout:    utils.GetDurationEnv("REDIS_WRITE_TIMEOUT", time.Second*3),
			DialTimeout:     utils.GetDurationEnv("REDIS_DIAL_TIMEOUT", time.Second*5),
		},
		Queue: &Queue{
			Concurrency: utils.GetIntEnv("QUEUE_WORKER_CONCURRENCY", 10),
		},
		Schedule: Schedule{
			Timezone: utils.GetEnv("SCHEDULE_TIMEZONE", "UTC"),
		},
	}

	return &cfg
}

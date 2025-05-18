package config

import (
	"github.com/robfig/cron/v3"
	"time"
)

type Config struct {
	App           *App
	Database      *Database
	SMTP          *SMTP
	DeliveryHttp  *DeliveryHttp
	DeliveryQueue *DeliveryQueue
	Redis         *Redis
	Queue         *Queue
	cron.Schedule
}

// App only this struct can deliver to module
type App struct {
	Name        string
	Environment string
	Version     string
	Mail        *Mail
}

type Mail struct {
	From string
}

type DeliveryHttp struct {
	Prefork       bool
	CaseSensitive bool
	StrictRouting bool
	BodyLimit     int
	ServerHeader  string
}

type DeliveryQueue struct {
	Concurrency int
}

type Database struct {
	User            string
	Password        string
	Host            string
	Port            int
	Database        string
	Params          string // opsional: tambahan param seperti charset=utf8mb4
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type SMTP struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type Redis struct {
	Addr            string
	Password        string
	DB              int
	PoolSize        int
	MinIdleConns    int
	ConnMaxLifetime time.Duration
	PoolTimeout     time.Duration
	ConnMaxIdleTime time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	DialTimeout     time.Duration
}

type Queue struct {
	Concurrency int
}

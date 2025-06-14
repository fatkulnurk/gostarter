package config

import (
	"os"
	"time"
)

type Config struct {
	App           *App
	Database      *Database
	DeliveryHttp  *DeliveryHttp
	DeliveryQueue *DeliveryQueue
	Redis         *Redis
	Queue         *Queue
	Schedule      *Schedule
	SMTP          *SMTP
}

// App only this struct can deliver to module
type App struct {
	Name        string
	Environment string
	Version     string
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

type Schedule struct {
	Timezone string
}

type SMTP struct {
	Host              string
	Port              int
	Username          string
	Password          string
	AuthType          string // one of => CRAM-MD5, CUSTOM, LOGIN, LOGIN-NOENC, NOAUTH, PLAIN, PLAIN-NOENC, XOAUTH2, SCRAM-SHA-1, SCRAM-SHA-1-PLUS, SCRAM-SHA-256, SCRAM-SHA-256-PLUS, SCRAM-SHA-384, SCRAM-SHA-384-PLUS, SCRAM-SHA-512, SCRAM-SHA-512-PLUS, AUTODISCOVER
	WithTLSPortPolicy int    // one of => 0 = Mandatory, 1 = Opportunistic, 2 = no tls
}

type SES struct {
	Region string
}

type S3 struct {
	Region               string
	Bucket               string
	AccessKey            string
	SecretKey            string
	Session              string
	Url                  string // url for generate url, if fill this field, it will be used to generate url for file, example https://minio.example.com for usePathStyleEndpoint = true, and https://bucket.minio.example.com for usePathStyleEndpoint = false
	UseStylePathEndpoint bool   // if true, format will be s3.amazonaws.com/bucket, if false, format will be bucket.s3.amazonaws.com
}

type LocalStorage struct {
	BasePath              string
	BaseURL               string
	DefaultDirPermission  os.FileMode // default 0755
	DefaultFilePermission os.FileMode // default 0644
}

package config

import (
	"github.com/fatkulnurk/gostarter/shared/constant"
	"github.com/joho/godotenv"
	"os"
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

	var cfg = Config{
		App: &App{
			Environment: env,
			Name:        os.Getenv("APP_NAME"),
			Version:     os.Getenv("APP_VERSION"),
			Mail: &Mail{
				From: os.Getenv("MAIL_FROM"),
			},
		},
		Database: &Database{
			User:            "",
			Password:        "",
			Host:            "",
			Port:            0,
			Database:        "",
			Params:          "",
			MaxOpenConns:    0,
			MaxIdleConns:    0,
			ConnMaxLifetime: 0,
			ConnMaxIdleTime: 0,
		},
		SMTP: &SMTP{
			From:     os.Getenv("SMTP_FROM"),
			Host:     os.Getenv("SMTP_HOST"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Port:     587,
			User:     os.Getenv("SMTP_USER"),
		},
	}

	return &cfg
}

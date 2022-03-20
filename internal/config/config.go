package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/todd-sudo/todo/pkg/logger"
)

type Config struct {
	Database struct {
		Host     string `yaml:"postgres_host" env:"POSTGRES_HOST"`
		User     string `yaml:"postgres_user" env:"POSTGRES_USER"`
		Password string `yaml:"postgres_password" env:"POSTGRES_PASSWORD"`
		DBName   string `yaml:"postgres_db" env:"POSTGRES_DB"`
		SslMode  string `yaml:"postgres_ssl" env:"POSTGRES_SSL_MODE"`
		Port     string `yaml:"postgres_db_port" env:"POSTGRES_PORT"`
	}
	App struct {
		Port string `yaml:"port"` // env-required:"true"
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}

package config

import "os"

type Config struct {
	Database struct {
		Host     string
		User     string
		Password string
		DBName   string
		SslMode  string
		Port     string
	}
	App struct {
		Port    string
		GinMode string
	}
}

// Конфигурация приложения
func GetConfig() *Config {

	return &Config{
		Database: struct {
			Host     string
			User     string
			Password string
			DBName   string
			SslMode  string
			Port     string
		}{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SslMode:  os.Getenv("POSTGRES_SSL_MODE"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
		App: struct {
			Port    string
			GinMode string
		}{
			Port:    os.Getenv("APP_PORT"),
			GinMode: os.Getenv("GIN_MODE"),
		},
	}
}

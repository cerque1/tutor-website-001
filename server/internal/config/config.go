package config

import "os"

type Config struct {
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string

	SecretKey string

	HTTPAddr string
}

func Load() *Config {
	return &Config{
		PostgresDB: os.Getenv("POSTGRES_DB"),
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost: os.Getenv("POSTGRES_HOST"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		HTTPAddr: os.Getenv("HTTP_ADDR"),
	}
}

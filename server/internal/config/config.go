package config

import "os"

type Config struct {
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string

	HTTPAddr string
}

func Load() *Config {
	return &Config{
		PostgresDB: os.Getenv("PostgresDB"),
		PostgresUser: os.Getenv("PostgresUser"),
		PostgresPassword: os.Getenv("PostgresPassword"),
		PostgresHost: os.Getenv("PostgresHost"),
		PostgresPort: os.Getenv("PostgresPort"),
		HTTPAddr: os.Getenv("HTTPAddr"),
	}
}

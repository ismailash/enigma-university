package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
		ApiHost: os.Getenv("API_HOST"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" {
		return errors.New("all environment variables required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

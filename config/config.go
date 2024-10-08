package config

import (
	"fmt"
	"os"

	"e-wallet-app/utils/common"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) ReadConfig() error {
	err := common.LoadEnv()
	if err != nil {
		return err
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" || c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" {
		return fmt.Errorf("Missing Required environment variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

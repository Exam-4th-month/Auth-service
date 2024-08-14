package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
		JWT      JWTConfig
	}
	JWTConfig struct {
		SecretKey string
	}
	ServerConfig struct {
		Port string
	}
	DatabaseConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.Server.Port = ":" + os.Getenv("SERVER_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")
	c.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")

	return nil
}

func New() (*Config, error) {
	var config Config
	if err := config.Load(); err != nil {
		return nil, err
	}
	return &config, nil
}

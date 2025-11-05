package helper

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	jwtSecretKey   string
	databaseConfig DatabaseConfig
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	AppConfig = &Config{
		jwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
		databaseConfig: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
		},
	}

	if AppConfig.jwtSecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is required")
	}

	log.Println("Configuration loaded successfully")
}

func (c *Config) GetJWTSecret() string {
	return c.jwtSecretKey
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return c.databaseConfig
}

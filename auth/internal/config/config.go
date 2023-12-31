package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	DBHost     string `yaml:"db_host" env:"DB_HOST" env-default:"localhost"`
	DBPort     string `yaml:"db_port" env:"DB_PORT" env-default:"27017"`
	DBUser     string `yaml:"db_user" env:"DB_USER" env-default:""`
	DBPassword string `yaml:"db_password" env:"DB_PASSWORD" env-default:""`
	DBName     string `yaml:"db_name" env:"DB_NAME" env-default:"cloud-drive"`
	HttpServer `yaml:"http_server"`
	RabbitMQ   `yaml:"rabbitmq"`
}

type HttpServer struct {
	GatewayURL string `yaml:"gateway_url" env:"GATEWAY_URL" env-default:"http://localhost:3000"`
	Host       string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port       string `yaml:"port" env:"PORT" env-default:"3001"`
}

type RabbitMQ struct {
	User     string `yaml:"user" env:"RABBITMQ_USER" env-default:"guest"`
	Password string `yaml:"password" env:"RABBITMQ_PASSWORD" env-default:"guest"`
	Host     string `yaml:"host" env:"RABBITMQ_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"RABBITMQ_PORT" env-default:"5672"`
	Queue    string `yaml:"queue" env:"RABBITMQ_QUEUE" env-default:"emails"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		log.Fatalf("error while reading config file: %s", err)
	}

	return &cfg
}

package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	DBHost     string `yaml:"db_host" env:"DB_HOST" env-default:"localhost"`
	DBPort     string `yaml:"db_port" env:"DB_PORT" env-default:"5432"`
	DBUser     string `yaml:"db_user" env:"DB_USER" env-default:""`
	DBPassword string `yaml:"db_password" env:"DB_PASSWORD" env-default:""`
	DBName     string `yaml:"db_name" env:"DB_NAME" env-default:"cloud-drive"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"PORT" env-default:"3000"`
}

func LoadConfig() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config.yml", &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}

package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	DB         `yaml:"db"`
	AWS        `yaml:"aws"`
}

type HTTPServer struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"PORT" env-default:"3000"`
}

type DB struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	DBName   string `yaml:"db_name" env-default:"files"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disabled"`
}

type AWS struct {
	Endpoint string `yaml:"endpoint"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSL      bool   `yaml:"ssl"`
}

func LoadConfig() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config.yml", &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}

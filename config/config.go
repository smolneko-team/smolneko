package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App     `yaml:"app"`
	HTTP    `yaml:"http"`
	Log     `yaml:"logger"`
	DB      `yaml:"postgres"`
	Storage `yaml:"storage"`
}

type App struct {
	Name        string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version     string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	StageStatus string `env-required:"true" yaml:"stage_status" env:"STAGE_STATUS"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Log struct {
	Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
}

type DB struct {
	PoolMax  int    `env-required:"true" yaml:"pg_pool_max" env:"DB_POOL_MAX"`
	Host     string `env-required:"true" yaml:"host" env:"DB_HOST"`
	Port     int    `env-required:"true" yaml:"port" env:"DB_PORT"`
	User     string `env-required:"true" yaml:"user" env:"DB_USER"`
	Password string `env-required:"true" yaml:"password" env:"DB_PASSWORD"`
	Name     string `env-required:"true" yaml:"name" env:"DB_NAME"`
	SSLMode  string `env-required:"true" yaml:"ssl_mode" env:"DB_SSL_MODE"`
}

type Storage struct {
	ImgKey  string `env-required:"true" yaml:"img_key" env:"IMG_KEY"`
	ImgSalt string `env-required:"true" yaml:"img_salt" env:"IMG_SALT"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	fmt.Printf("%+v", cfg)

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config -
type Config struct {
	App     `yaml:"app"`
	HTTP    `yaml:"http"`
	Log     `yaml:"logger"`
	DB      `yaml:"postgres"`
	Storage `yaml:"storage"`
}

// App -
type App struct {
	Name        string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version     string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	StageStatus string `env-required:"true" yaml:"stage_status" env:"STAGE_STATUS"`
}

// HTTP -
type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

// Log -
type Log struct {
	Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
}

// DB struct with PostgreSQL configuration.
type DB struct {
	Host     string `env-required:"true" yaml:"host" env:"DB_HOST"`
	Port     int    `env-required:"true" yaml:"port" env:"DB_PORT"`
	User     string `env-required:"true" yaml:"user" env:"DB_USER"`
	Password string `env-required:"true" yaml:"password" env:"DB_PASSWORD"`
	Name     string `env-required:"true" yaml:"name" env:"DB_NAME"`
	SSLMode  string `env-required:"true" yaml:"ssl_mode" env:"DB_SSL_MODE"`
	PoolMax  int    `env-required:"true" yaml:"pg_pool_max" env:"DB_POOL_MAX"`
}

// Storage config struct for S3 image proxying with ImgProxy.
//
// About serving files from S3 — https://docs.imgproxy.net/serving_files_from_s3
//
// Configuration — https://docs.imgproxy.net/configuration
//
// Processing options — https://docs.imgproxy.net/generating_the_url?id=processing-options
type Storage struct {
	ImgKey   string `env-required:"true" yaml:"img_key" env:"IMG_KEY"`
	ImgSalt  string `env-required:"true" yaml:"img_salt" env:"IMG_SALT"`
	Bucket   string `env-required:"true" yaml:"bucket" env:"BUCKET"`
	ImgURL   string `env-required:"true" yaml:"img_url" env:"IMG_URL"`
	ProcOpts string `env-required:"true" yaml:"proc_opts" env:"PROC_OPTS"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

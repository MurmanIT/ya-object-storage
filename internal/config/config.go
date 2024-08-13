package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type S3 struct {
	Region    string `yaml:"region" env-required:"true" env-default:"ru-central1"`
	Bucket    string `yaml:"bucket" env-required:"true" env:"S3_BUCKET"`
	Url       string `yaml:"url" env-required:"true" env-default:"https://storage.yandexcloud.net"`
	AccessKey string `yaml:"access_key" env-required:"true" env:"S3_ACCESS_KEY"`
	SecretKey string `yaml:"secret_key" env-required:"true" env:"S3_SECRET_KEY"`
}

type HttpServer struct {
	Port int    `yaml:"port" env-required:"true" env-default:"8080"`
	User string `yaml:"user" env-required:"true"`
	Pass string `yaml:"password" env-required:"true" env: "HTTP_SERVER_PASSWORD"`
}

type UploadDir struct {
	Dir string `yaml:"dir" env-required:"true" env-default:"./files" env:"UPLOAD_FILES_DIR"`
}

type Config struct {
	Env        string     `yaml:"env" env-default:"dev" env-required:"true"`
	S3         S3         `yaml:"s3"`
	HttpServer HttpServer `yaml:"http_server"`
	UploadDir  UploadDir
}

func LoadConfig() (*Config, error) {
	var cfg Config

	configPath := getConfig(os.Getenv("CONFIG_PATH"))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
		return nil, err
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("config error: %s", err)
		return nil, err
	}

	return &cfg, nil
}

func getConfig(env string) string {
	if env == "" {
		log.Fatal("env is required")
	}
	return "./config/" + env + ".yaml"
}

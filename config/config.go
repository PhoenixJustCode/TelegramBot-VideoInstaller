package config

import (
    "github.com/kelseyhightower/envconfig"
)

type Config struct {
    TelegramToken string `envconfig:"TELEGRAM_TOKEN" required:"true"`
    MaxFileSizeMB int    `envconfig:"MAX_FILE_SIZE_MB" default:"50"`
    DownloadDir   string `envconfig:"DOWNLOAD_DIR" default:"/tmp/videos"`
    OutputDir     string `envconfig:"OUTPUT_DIR" default:"/tmp/audio"`
}

func LoadConfig() (*Config, error) {
    var cfg Config
    err := envconfig.Process("api", &cfg)
    return &cfg, err
}

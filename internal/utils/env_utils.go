package utils

import (
	"os"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type TokenCFG struct {
	Token         string
	MaxFileSizeMB int64
	DownloadDir   string
	OutputDir     string
}

func LoadFromEnv() TokenCFG {
	var cfg TokenCFG

	cfg.Token = os.Getenv("API_TOKEN")
	cfg.DownloadDir = os.Getenv("API_DOWNLOADDIR")
	cfg.OutputDir = os.Getenv("API_OUTPUTDIR")

	maxFileSizeStr := os.Getenv("API_MAXFILESIZEMB")
	maxFileSize, err := strconv.ParseInt(maxFileSizeStr, 10, 64)
	if err != nil {
		log.Warn("Некорректное значение API_MAXFILESIZEMB, установлено значение по умолчанию 100")
		maxFileSize = 100
	}
	cfg.MaxFileSizeMB = maxFileSize

	return cfg
}

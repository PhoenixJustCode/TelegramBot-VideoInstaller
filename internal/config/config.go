package config

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
	cfg := TokenCFG{
		Token:       os.Getenv("API_TOKEN"),
		DownloadDir: os.Getenv("API_DOWNLOADDIR"),
		OutputDir:   os.Getenv("API_OUTPUTDIR"),
	}

	maxFileSizeStr := os.Getenv("API_MAXFILESIZEMB")
	maxFileSize, err := strconv.ParseInt(maxFileSizeStr, 10, 64)
	if err != nil || maxFileSize <= 0 {
		log.Warn("⚠️ Invalid API_MAXFILESIZEMB value, defaulting to 100 MB")
		maxFileSize = 100
	}
	cfg.MaxFileSizeMB = maxFileSize

	return cfg
}

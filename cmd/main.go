package main

import (
	"os"
	"telegramBotInstaller/internal/bot"
	"telegramBotInstaller/internal/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Формат логов: можно JSON или красивый текст
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true, 
	})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)

	if err := godotenv.Load(); err != nil {
		log.Warnf("❗ Не удалось загрузить .env файл: %v", err)
	}
}


func main() {
	cfg := config.LoadFromEnv()
	bot.StartBot(cfg)
}

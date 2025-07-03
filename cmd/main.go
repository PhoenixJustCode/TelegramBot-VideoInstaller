package main

import (
	"os"
	"telegramBotInstaller/internal/bot"
	"telegramBotInstaller/internal/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)

	err := godotenv.Load()
	if err != nil {
		log.Warn("Не удалось загрузить .env файл:", err)
	}
}

func main() {
	cfg := utils.LoadFromEnv()
	bot.StartBot(cfg)
}

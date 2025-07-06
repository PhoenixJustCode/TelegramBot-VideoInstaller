package bot

import (
	"strings"
	"telegramBotInstaller/internal/config"
	"telegramBotInstaller/internal/services"
	"telegramBotInstaller/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func StartBot(cfg config.TokenCFG) {
	if cfg.Token == "" {
		log.Fatal("Telegram API token is not set. Check environment variables.")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("✅ Authorized as %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "/help" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, GetHelpMessage())
				if _, err := bot.Send(msg); err != nil {
					log.Printf("❌ Ошибка при отправке сообщения: %v", err)
				}
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
				if _, err := bot.Request(callback); err != nil {
					log.Printf("❌ Ошибка при выполнении callback: %v", err)
				}
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		username := update.Message.From.FirstName + " " + update.Message.From.LastName
		log.Infof("[%s] %s", username, update.Message.Text)

		userText := update.Message.Text

		if update.Message.Video != nil && update.Message.Video.MimeType == "video/mp4" {
			videoFileID := update.Message.Video.FileID

			videoPath, err := services.DownloadVideo(bot, videoFileID, cfg.DownloadDir)
			if err != nil {
				SendMessage(update, "❌ Download error", bot)
				log.Fatal(err)
				continue
			}

			ProcessAndSendAudio(bot, update, videoPath, videoFileID, cfg.OutputDir)
			continue
		}

		if utils.IsYoutubeURL(userText) {
			videoID, err := utils.ExtractVideoID(userText)
			if err != nil {
				SendMessage(update, "❌ Это не ссылка на видео YouTube", bot)
				continue
			}

			SendMessage(update, "📺 YouTube ссылка получена! Скачиваю...", bot)

			videoPath, err := services.DownloadYouTubeVideo(userText, cfg.DownloadDir)
			if err != nil {
				SendMessage(update, "❌ Ошибка при скачивании YouTube видео", bot)
				log.Printf("❌ Ошибка загрузки: %v", err)
				continue
			}

			ProcessAndSendAudio(bot, update, videoPath, videoID, cfg.OutputDir)
			continue
		}

		userTextLower := strings.ToLower(userText)
		switch userTextLower {
		case "/start":
			reply := "Hello, " + username + "! Welcome 👋\n\nTo see what I can do, press the button below or type /help"
			helpButton := tgbotapi.NewInlineKeyboardButtonData("🆘 Help", "/help")
			SendMessageWithKeyboard(update, reply, bot, helpButton)

		case "/help":
			SendMessage(update, GetHelpMessage(), bot)

		default:
			reply := "You wrote: " + update.Message.Text
			SendMessage(update, reply, bot)
		}
	}
}

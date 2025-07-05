package bot

import (
	"path/filepath"
	"strings"
	"telegramBotInstaller/internal/config"
	"telegramBotInstaller/internal/services"
	"telegramBotInstaller/internal/utils"
	"time"

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
				bot.Send(msg)
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
				bot.Request(callback)
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		username := update.Message.From.FirstName + " " + update.Message.From.LastName
		log.Infof("[%s] %s", username, update.Message.Text)

		userText := strings.ToLower(update.Message.Text)

		if update.Message.Video != nil && update.Message.Video.MimeType == "video/mp4" {
			videoFileID := update.Message.Video.FileID

			videoPath, err := services.DownloadVideo(bot, videoFileID, cfg.DownloadDir)
			if err != nil {
				SendMessage(update, "❌ Download error", bot)
				log.Fatal(err)
				continue
			}

			_, err = services.ConvertMp4ToMp3WithID(videoPath, cfg.OutputDir, videoFileID)
			if err != nil {
				SendMessage(update, "❌ Conversion error", bot)
				log.Fatal(err)
				continue
			}

			audioPath := filepath.Join(cfg.OutputDir, videoFileID+".mp3")
			audioFile := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FilePath(audioPath))
			bot.Send(audioFile)
			time.Sleep(time.Second*10)
			
			if err := utils.DeleteFile(videoPath); err != nil {
				log.Errorf("❌ Failed to delete video file: %v", err)
			} else {
				log.Infof("✅ Video file deleted: %s", videoPath)
			}

			if err := utils.DeleteFile(audioPath); err != nil {
				log.Errorf("❌ Failed to delete MP3 file: %v", err)
			} else {
				log.Infof("✅ MP3 file deleted: %s", videoPath)
			}
			continue
		}

		if utils.IsYoutubeURL(userText) {
			reply := "📺 Got YouTube link! Processing..."
			SendMessage(update, reply, bot)
			continue
		}

		switch userText {
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

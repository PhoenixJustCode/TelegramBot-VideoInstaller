package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"telegramBotInstaller/internal/services"
	"telegramBotInstaller/internal/utils"
	"time"
)

func SendMessageWithKeyboard(update tgbotapi.Update, reply string, bot *tgbotapi.BotAPI, helpButton tgbotapi.InlineKeyboardButton) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(helpButton),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		log.Printf("❌ Ошибка при отправке сообщения: %v", err)
	}
}

func SendMessage(update tgbotapi.Update, reply string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	msg.ReplyToMessageID = update.Message.MessageID
	if _, err := bot.Send(msg); err != nil {
		log.Printf("❌ Ошибка при отправке сообщения: %v", err)
	}
}

func GetHelpMessage() string {
	return `👋 Hello! I am a bot for video and audio conversion 🎧

What I can do:
✅ Send me a video in MP4 format — I will convert it to MP3 and send it back to you.
✅ Send a YouTube link — I will download it and convert to MP3 too.

Use the /start command to begin or just send a video/link.

❓ If you have any questions — type /help anytime.

Enjoy! 🚀`
}

func ProcessAndSendAudio(bot *tgbotapi.BotAPI, update tgbotapi.Update, videoPath, audioID, outputDir string) {
	_, err := services.ConvertMp4ToMp3WithID(videoPath, outputDir, audioID)
	if err != nil {
		SendMessage(update, "❌ Ошибка при конвертации", bot)
		log.Fatal(err)
		return
	}

	audioPath := filepath.Join(outputDir, audioID+".mp3")
	audioFile := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FilePath(audioPath))
	if _, err := bot.Send(audioFile); err != nil {
		log.Printf("❌ Ошибка при отправке MP3: %v", err)
	}

	time.Sleep(time.Second * 10)

	if err := utils.DeleteFile(videoPath); err != nil {
		log.Errorf("❌ Не удалось удалить видео: %v", err)
	} else {
		log.Infof("✅ Видео удалено: %s", videoPath)
	}

	if err := utils.DeleteFile(audioPath); err != nil {
		log.Errorf("❌ Не удалось удалить MP3: %v", err)
	} else {
		log.Infof("✅ MP3 удалено: %s", audioPath)
	}
}

package bot

import (
	log "github.com/sirupsen/logrus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

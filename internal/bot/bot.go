package bot

import (
	"strings"
	"telegramBotInstaller/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func StartBot(cfg utils.TokenCFG) {
	if cfg.Token == "" {
		log.Fatal("Токен Telegram API не задан. Проверьте переменные окружения.")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {	// 1. Обработка CallbackQuery (кнопки)
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "/help" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, getHelpMessage())
				bot.Send(msg)

				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "") // Ответ на callback, чтобы убрать «часики»
				bot.Request(callback)
			}
			continue
		}

		if update.Message == nil { //2. Обработка обычных сообщений
			continue
		}

		username := update.Message.From.FirstName + " " + update.Message.From.LastName
		userText := strings.ToLower(update.Message.Text)

		log.Infof("[%s] %s", username, update.Message.Text)

		switch userText {
		case "/start":
			reply := "Привет, " + username + "! Добро пожаловать 👋\n\nЧтобы узнать, что я умею — нажми кнопку ниже или напиши /help"
			helpButton := tgbotapi.NewInlineKeyboardButtonData("🆘 Помощь", "/help")
			sendMessageWithKeyboard(update, reply, bot, helpButton)

		case "/help":
			sendMessage(update, getHelpMessage(), bot)

		default:
			reply := "Вы написали: " + update.Message.Text
			sendMessage(update, reply, bot)
		}
	}
}

func sendMessage(update tgbotapi.Update, reply string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func sendMessageWithKeyboard(update tgbotapi.Update, reply string, bot *tgbotapi.BotAPI, helpButton tgbotapi.InlineKeyboardButton) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(helpButton),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func getHelpMessage() string {
	return `👋 Привет! Я бот для конвертации видео и аудио 🎧

Что я умею:
✅ Отправь мне видео в формате MP4 — я преобразую его в аудиофайл MP3 и пришлю тебе обратно.
✅ Скинь ссылку на видео с YouTube — я загружу его и тоже конвертирую в MP3.

Используй команду /start чтобы начать или просто пришли видео/ссылку.

❓ Если возникли вопросы — пиши /help в любое время.

Приятного использования! 🚀`
}

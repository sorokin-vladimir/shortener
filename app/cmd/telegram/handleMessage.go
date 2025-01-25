package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sorokin-vladimir/shortener/internal/shortener"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	texts := strings.Split(strings.ToLower(update.Message.Text), " ")
	url, short, expiry := utils.GetArgs(strings.Join(texts, " "))
	var msg tgbotapi.MessageConfig

	if !strings.HasPrefix(url, "http") {
		log.Printf("Did not provide an URL in Message | User: %d", update.Message.From.ID)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "It doesn't look like a valid URL. Please provide a valid URL.")
		bot.Send(msg)
		return
	}

	result := shortener.Shortener(
		url,
		"",
		update.Message.From.ID,
		expiry,
		short,
	)

	if result.Err != nil {
		log.Println(result.Err)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, result.Err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, result.Short)
	}

	bot.Send(msg)
}

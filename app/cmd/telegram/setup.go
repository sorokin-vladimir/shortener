package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func setupBotCommands(bot *tgbotapi.BotAPI) {
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Start bot"},
		{Command: "help", Description: "Commands list"},
		{Command: "short", Description: "/short <URL> <custom short> <custom expiry> -> Get shorten url"},
	}
	_, err := bot.Request(tgbotapi.NewSetMyCommands(commands...))
	if err != nil {
		log.Panic(err)
	}
	log.Println("Commands setup completed")
}

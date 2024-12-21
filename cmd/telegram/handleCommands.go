package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sorokin-vladimir/shortener/internal/shortener"
	"github.com/sorokin-vladimir/shortener/internal/utils"
)

func handleCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi! I'm URL-Shortener bot. Use /short <URL> to create short URL alias or /help to get to know all commands")
		bot.Send(msg)

	case "help":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Available commands:\n/start - Run the bot\n/help - Commands list\n/short - /short <URL> <custom short> <custom expiry> - Get shorten url")
		bot.Send(msg)

	case "short":
		url, short, expiry := utils.GetArgs(update.Message.CommandArguments())
		var msg tgbotapi.MessageConfig

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

	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Use /help to search commands.")
		bot.Send(msg)
	}
}

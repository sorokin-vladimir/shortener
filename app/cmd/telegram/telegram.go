package telegram

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(ctx context.Context) error {
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		log.Fatal("TELEGRAM_API_KEY is not set up")
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return err
	}

	log.Printf("Telegram bot is authorized as %s", bot.Self.UserName)

	setupBotCommands(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Message handling
	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				if update.Message.IsCommand() {
					handleCommands(bot, update)
				} else if update.Message.Text != "" {
					handleMessage(bot, update)
				}
			}
		case <-ctx.Done():
			log.Println("Stopping Telegram bot...")
			return nil
		}
	}
}

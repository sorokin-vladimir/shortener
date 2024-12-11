package telegram

import (
	"context"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
)

func StartBot(ctx context.Context, rdb *redis.Client) error {
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		log.Fatal("TELEGRAM_API_KEY is not set up")
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return err
	}

	log.Printf("Telegram bot is authorized as %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Message handling
	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				handleMessage(update.Message, bot, rdb)
			}
		case <-ctx.Done():
			log.Println("Stopping Telegram bot...")
			return nil
		}
	}
}

func handleMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI, rdb *redis.Client) {
	// Example message handling and working with Redis
	text := msg.Text
	log.Printf("[%s] %s", msg.From.UserName, text)

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Your message: "+text)
	bot.Send(reply)
}

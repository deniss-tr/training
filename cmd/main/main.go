package main

import (
	"log"
	"os"

	"github.com/deniss-tr/training/internal/app/commands"
	"github.com/deniss-tr/training/internal/service/day"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	dayService := day.NewService()

	commander := commands.NewCommander(bot, dayService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}

package commands

import (
	"github.com/deniss-tr/training/internal/service/day"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot        *tgbotapi.BotAPI
	dayService *day.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	dayService *day.Service,
) *Commander {
	return &Commander{
		bot:        bot,
		dayService: dayService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		default:
			c.Default(update.Message)
		}

	}
}

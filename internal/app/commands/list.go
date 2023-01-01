package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	days := c.dayService.List()
	outputText := ""

	for _, d := range days {
		outputText += d.Title
		outputText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputText)
	c.bot.Send(msg)
}

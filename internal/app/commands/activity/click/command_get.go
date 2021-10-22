package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *clickCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("clickCommander.Get: error parsing product index: %v\n", err)

		c.SendMessageToChat(
			tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error parsing product index: %v", err)),
			"clickCommander.Get",
		)

		return
	}

	m, err := c.service.Describe(idx)
	if err != nil {
		log.Printf("clickCommander.Get: error getting product by index %d: %v\n", idx, err)

		c.SendMessageToChat(
			tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error: %s", err.Error())),
			"clickCommander.Get",
		)

		return
	}

	c.SendMessageToChat(tgbotapi.NewMessage(inputMsg.Chat.ID, m.String()), "clickCommander.Get")
}

package click

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
)

func (c *clickCommander) New(inputMsg *tgbotapi.Message) {
	var m activity.Click

	args := inputMsg.CommandArguments()

	err := json.Unmarshal([]byte(args), &m)
	if err != nil {
		outputMsgText := "Error parsing model's data. To add new item please enter the data in format {\"title\":\"myItem\"}"
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancel", fmt.Sprintf("%s__saveNewItem__cancel", ActivityClickPrefix)),
			),
		)

		msg.ReplyMarkup = keyboard

		c.SendMessageToChat(msg, "clickCommander.New")

		return
	}

	idx, err := c.service.Create(m)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error creating new item: %v", err))

		c.SendMessageToChat(msg, "clickCommander.New")

		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Item succesfully created. Index: %d", idx))

	c.SendMessageToChat(msg, "clickCommander.New")
}

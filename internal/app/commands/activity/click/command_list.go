package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *clickCommander) List(inputMsg *tgbotapi.Message) {
	pagedList, err := c.service.List(c.cursor, c.limit)
	outputMsgText := ""

	if err != nil {
		outputMsgText := "Oops, something went wrong: " + err.Error()

		c.SendMessageToChat(tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText), "clickCommander.List")

		return
	}

	if len(pagedList) == 0 {
		outputMsgText = fmt.Sprintf(
			"List of models is empty. You can use /new__%s command to add new items to it",
			ActivityClickPrefix,
		)

		c.SendMessageToChat(tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText), "clickCommander.List")

		return
	}

	outputMsgText = "List of items:\n\n"

	i := c.cursor
	for _, item := range pagedList {
		outputMsgText += fmt.Sprintf("%d: %s\n", i, item.String())
		i++
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if c.cursor > 0 {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Previous Page", fmt.Sprintf("%s__list__prevPage", ActivityClickPrefix)),
		)
	}

	list, err := c.service.List(0, 0)
	if err != nil {
		outputMsgText := "Oops, something went wrong: " + err.Error()

		c.SendMessageToChat(tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText), "clickCommander.List")

		return
	}

	total := len(list)
	if int(c.cursor+c.limit) < total {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next Page", fmt.Sprintf("%s__list__nextPage", ActivityClickPrefix)),
		)
	}

	if len(buttons) > 0 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))

		msg.ReplyMarkup = keyboard
	}

	c.SendMessageToChat(msg, "clickCommander.List")
}

package click

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
	"log"
	"strconv"
	"strings"
)

func (c *clickCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	parts := strings.Split(args, " ")

	if len(parts) != 2 {
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Unable to parse arguments %s. Please use format: `item_index {\"title\":\"myItem\"}`", args),
		)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancel", fmt.Sprintf("%s__editItem__cancel", ActivityClickPrefix)),
			),
		)

		msg.ReplyMarkup = keyboard
		msg.ParseMode = "markdown"

		c.SendMessageToChat(msg, "clickCommander.Edit")
		return
	}

	idx, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		log.Printf("clickCommander.Edit: error parsing product index: %v\n", err)

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error parsing product index: %v", err))

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancel", fmt.Sprintf("%s__editItem__cancel", ActivityClickPrefix)),
			),
		)

		msg.ReplyMarkup = keyboard

		c.SendMessageToChat(msg, "clickCommander.Edit")

		return
	}

	var m activity.Click

	err = json.Unmarshal([]byte(parts[1]), &m)
	if err != nil {
		outputMsgText := "Error parsing item's data. To edit item please enter the data in format `item_index {\"title\":\"myItem\"}`"
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancel", fmt.Sprintf("%s__editItem__cancel", ActivityClickPrefix)),
			),
		)

		msg.ReplyMarkup = keyboard
		msg.ParseMode = "markdown"

		c.SendMessageToChat(msg, "clickCommander.Edit")

		return
	}

	err = c.service.Update(idx, m)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Error updating item: %v", err))

		c.SendMessageToChat(msg, "clickCommander.Edit")

		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Item %d succesfully updated", idx))

	c.SendMessageToChat(msg, "clickCommander.Edit")
}

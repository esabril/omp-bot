package click

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *clickCommander) Help(inputMsg *tgbotapi.Message) {
	outputMessageText := fmt.Sprintf(`List of commands:

	/help__%[1]s — shows this help
	/list__%[1]s — show items list
	/get__%[1]s — get item by index
	<b>Format:</b> /get__%[1]s 1

	/new__%[1]s — add new item to list
	<b>Format:</b> /new__%[1]s {"title":"myItem"}

	/edit__%[1]s — edit item in list
	<b>Format:</b> /edit__%[1]s 1 {"title":"myItem"}

	/delete__%[1]s — delete item from list
	<b>Format:</b> /delete__%[1]s 1`,
		ActivityClickPrefix,
	)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMessageText)
	msg.ParseMode = "html"

	c.SendMessageToChat(msg, "clickCommander.Help")
}

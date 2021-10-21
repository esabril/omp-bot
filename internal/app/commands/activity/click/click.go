package click

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/activity/click"
	"log"
)

const ActivityClickPrefix = "activity__click"

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ClickCommander interface {
	Commander
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type clickCommander struct {
	bot     *tgbotapi.BotAPI
	service service.ClickService
	cursor  uint64
	limit   uint64
}

func NewActivityClickCommander(bot *tgbotapi.BotAPI, service service.ClickService) ClickCommander {
	return &clickCommander{
		bot:     bot,
		service: service,
		cursor:  0,
		limit:   1, // small value only for pagination's demonstration
	}
}

func (c *clickCommander) SendMessageToChat(message tgbotapi.MessageConfig, sendingErrorLogCase string) {
	_, err := c.bot.Send(message)

	if err != nil {
		log.Printf("%s: error sending message to chat: %v", sendingErrorLogCase, err)
	}
}

func (c *clickCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "saveNewItem":
		c.SaveNewItem(callback, callbackPath)
	case "editItem":
		c.EditItem(callback, callbackPath)
	case "list":
		c.Paginate(callback, callbackPath)
	}
}

func (c *clickCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.FlushCursor()
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		log.Printf("clickCommander.HandleCommand: unknown command - %s", commandPath.CommandName)
	}
}

func (c *clickCommander) SaveNewItem(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "cancel":
		c.SendMessageToChat(
			tgbotapi.NewMessage(callback.Message.Chat.ID, "Сreation of a new item canceled"),
			"clickCommander.HandleCommand.SaveNewItem",
		)
	default:
		log.Printf("clickCommander.HandleCommand.SaveNewItem: unknown command - %s", callbackPath.CallbackData)
	}
}

func (c *clickCommander) EditItem(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "cancel":
		c.SendMessageToChat(
			tgbotapi.NewMessage(callback.Message.Chat.ID, "Item update canceled"),
			"clickCommander.HandleCommand.EditItem",
		)
	default:
		log.Printf("clickCommander.HandleCommand.EditItem: unknown command - %s", callbackPath.CallbackData)
	}
}

func (c *clickCommander) ModifyCursorsForNextPage() {
	c.cursor += c.limit
}

func (c *clickCommander) ModifyCursorsForPrevPage() {
	c.cursor -= c.limit
}

func (c *clickCommander) FlushCursor() {
	c.cursor = 0
}

func (c *clickCommander) Paginate(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "prevPage":
		c.ModifyCursorsForPrevPage()
		c.List(callback.Message)
	case "nextPage":
		c.ModifyCursorsForNextPage()
		c.List(callback.Message)
	}
}

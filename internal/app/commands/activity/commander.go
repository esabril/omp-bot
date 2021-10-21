package activity

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/activity/click"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/activity/click"
	"log"
)

type ClickCommander struct {
	bot            *tgbotapi.BotAPI
	clickCommander click.ClickCommander
}

func NewClickCommander(bot *tgbotapi.BotAPI) *ClickCommander {
	s := service.NewDummyClickService()

	return &ClickCommander{
		bot:            bot,
		clickCommander: click.NewActivityClickCommander(bot, s),
	}
}

func (c *ClickCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "click":
		c.clickCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ClickCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ClickCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "click":
		c.clickCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ClickCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

package bot

import (
	"tg-pill-bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func New(config config.Config) (*Bot, error) {
	botApi, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}
	botApi.Debug = false
	return &Bot{api: botApi}, nil
}

func (b *Bot) Run() {
	for update := range b.api.GetUpdatesChan(tgbotapi.UpdateConfig{Offset: 0, Timeout: 60}) {
		if update.Message == nil {
			continue
		}
	}
}

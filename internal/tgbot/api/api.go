package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joomcode/errorx"
	"github.com/mih-kopylov/my-gorzdrav-spb-bot/internal/config"
	"go.uber.org/zap"
)

func NewTgbotAPI(logger *zap.Logger, conf *config.Config) (*tgbotapi.BotAPI, error) {
	err := tgbotapi.SetLogger(&TgLogger{logger: logger})
	if err != nil {
		return nil, errorx.EnhanceStackTrace(err, "failed to configure bot api logging")
	}

	api, err := tgbotapi.NewBotAPI(conf.TelegramApiToken)
	if err != nil {
		return nil, errorx.EnhanceStackTrace(err, "failed to create bot")
	}

	api.Debug = true
	return api, nil
}

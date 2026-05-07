package app

import (
	"context"

	"github.com/mih-kopylov/my-gorzdrav-spb-bot/internal/config"
	"github.com/mih-kopylov/my-gorzdrav-spb-bot/internal/tgbot"
	"github.com/mih-kopylov/my-gorzdrav-spb-bot/internal/tgbot/api"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func RunApplication() {
	fx.New(
		fx.Provide(NewLogger),
		fx.Provide(config.NewConfig),
		fx.WithLogger(convertLogger),
		fx.Provide(api.NewTgbotAPI),
		fx.Provide(tgbot.NewTgbot),
		fx.Invoke(Run),
	).Run()
}

func Run(bot *tgbot.TgBot) error {
	return bot.Start()
}

func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	result, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	lc.Append(
		fx.StopHook(
			func(_ context.Context) error {
				return result.Sync()
			},
		),
	)

	return result, nil
}

func convertLogger(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: logger}
}

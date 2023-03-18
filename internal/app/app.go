package app

import (
	"context"
	"github.com/mih-kopylov/my-gorzdrav-spb-bot/internal/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func RunApplication() {
	fx.New(
		fx.Provide(NewLogger),
		fx.Provide(config.NewConfig),
		fx.WithLogger(convertLogger),
		fx.Invoke(Run),
	).Run()
}

func Run(_ *config.Config, _ *zap.Logger) {
}

func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	result, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.StopHook(func(_ context.Context) error {
		return result.Sync()
	}))

	return result, nil
}

func convertLogger(logger *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: logger}
}

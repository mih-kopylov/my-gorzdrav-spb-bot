package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
)

type TgBot struct {
	logger *zap.Logger
	api    *tgbotapi.BotAPI
}

func NewTgbot(logger *zap.Logger, api *tgbotapi.BotAPI) (*TgBot, error) {
	return &TgBot{
		logger: logger,
		api:    api,
	}, nil
}

func (b *TgBot) Start() error {
	err := b.registerCommands()
	if err != nil {
		return err
	}

	go func() {
		b.processUpdates()
	}()

	return nil
}

func (b *TgBot) processUpdates() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 10

	updates := b.api.GetUpdatesChan(updateConfig)
	for update := range updates {
		err := b.callHandler(update)
		if err != nil {
			fromChat := update.FromChat()
			var chatId int64
			if fromChat != nil {
				chatId = fromChat.ID
			}
			b.logger.Error(
				"failed to handle update",
				zap.Int64("chat", chatId),
				zap.Error(errorx.EnsureStackTrace(err)),
			)
		}
	}
}

func (b *TgBot) callHandler(update tgbotapi.Update) error {
	switch {
	case update.Message != nil:
		return b.handleMessage(update.Message)
	default:
		return errorx.IllegalArgument.New("unsupported update type")
	}
}

func (b *TgBot) registerCommands() error {
	setMyCommandsConfig := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "start",
			Description: "Start the app",
		},
		tgbotapi.BotCommand{
			Command:     "example",
			Description: "Example command",
		},
	)
	_, err := b.api.Request(setMyCommandsConfig)
	if err != nil {
		return errorx.EnhanceStackTrace(err, "failed to register bot commands")
	}

	return nil
}

func (b *TgBot) handleMessage(message *tgbotapi.Message) error {
	commandName := message.Command()

	if commandName != "" {
		b.logger.Info(
			"command run",
			zap.Int64("chatId", message.Chat.ID),
			zap.String("command", commandName),
		)
	} else {
		b.logger.Info(
			"message received",
			zap.Int64("chatId", message.Chat.ID),
			zap.String("text", message.Text),
		)
	}

	return nil
}

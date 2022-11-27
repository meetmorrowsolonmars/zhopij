package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.DPanicLevel))
	if err != nil {
		fmt.Println("can not run the app because the logger is not created:", err)
		os.Exit(1)
	}

	sugared := logger.Sugar()
	defer func() {
		_ = sugared.Sync()
	}()

	const telegramApiToken = "ZHOPIJ_TELEGRAM_API_TOKEN"
	token := os.Getenv(telegramApiToken)
	if len(token) == 0 {
		sugared.Fatalln("can not run the app: telegram api token not provided")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		sugared.Fatalln("can not run the app:", err)
	}

	config := tgbotapi.NewUpdate(0)
	config.Timeout = 30
	updates := bot.GetUpdatesChan(config)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
		if err != nil {
			sugared.Errorw("can not send message", zap.Error(err), zap.Int64("chat_id", update.Message.Chat.ID))
		}
	}
}

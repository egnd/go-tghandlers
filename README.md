# Telegram chained events handler

[![PkgGoDev](https://pkg.go.dev/badge/github.com/egnd/go-tghandlers/v1)](https://pkg.go.dev/github.com/egnd/go-tghandlers/v1?tab=doc)
[![GoDoc](https://godoc.org/github.com/egnd/go-tghandlers?status.svg)](https://pkg.go.dev/github.com/egnd/go-tghandlers/v1)
[![Go Report Card](https://goreportcard.com/badge/github.com/egnd/go-tghandlers)](https://goreportcard.com/report/github.com/egnd/go-tghandlers)
[![codecov.io](https://codecov.io/github/egnd/go-tghandlers/coverage.svg?branch=master)](https://codecov.io/gh/egnd/go-tghandlers?branch=master)
[![Build](https://github.com/egnd/go-tghandlers/workflows/Build/badge.svg)](https://github.com/egnd/go-tghandlers/actions?query=branch%3Amaster)
[![Unit](https://github.com/egnd/go-tghandlers/workflows/Unit/badge.svg)](https://github.com/egnd/go-tghandlers/actions?query=branch%3Amaster)
[![Integration](https://github.com/egnd/go-tghandlers/workflows/Integration/badge.svg)](https://github.com/egnd/go-tghandlers/actions?query=branch%3Amaster)
[![API](https://github.com/egnd/go-tghandlers/workflows/API/badge.svg)](https://github.com/egnd/go-tghandlers/actions?query=branch%3Amaster)

### Example
```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/egnd/go-tghandlers/pkg/handlers"
	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func main() {
	// init services
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		log.Fatal(err)
		return
	}

	api, err := tgbotapi.NewBotAPIWithClient("BOT_TOKEN", &http.Client{})
	if err != nil {
		logger.Fatal("init tg api", zap.Error(err))
		return
	}

	updSettings := tgbotapi.NewUpdate(0)
	updSettings.Timeout = 60
	updChannel, err := api.GetUpdatesChan(updSettings)
	if err != nil {
		logger.Fatal("init updates channel", zap.Error(err))
		return
	}

	bot := listener.NewListener()

	// set handlers
	bot.Add(listener.EventPrivateMessage,
		handlers.AppendLogWithMessage,
		handlers.LogIncoming,
		func(next listener.EventHandler) listener.EventHandler {
			return func(ctx context.Context, update tgbotapi.Update, lggr listener.ILogger) {
				response := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, world!")
				response.ReplyToMessageID = update.Message.MessageID

				if _, err = api.Send(response); err != nil {
					lggr.Error("send response", zap.Error(err))
				}
			}
		},
	)

	// start listening
	if err = bot.Listen(context.Background(), updChannel, api, logger); err != nil {
		logger.Fatal("bot listening", zap.Error(err))
	}
}
```
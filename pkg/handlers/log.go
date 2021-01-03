// Package handlers is ...
package handlers

import (
	"context"

	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

func AppendLogWithMessage(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		next(ctx, update, logger.With(zap.String("tg.message", update.Message.Text)))
	}
}

func AppendLogWithInlineQuery(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		next(ctx, update, logger.With(zap.String("tg.inline_query", update.InlineQuery.Query)))
	}
}

func AppendLogWithCallbackQuery(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		next(ctx, update, logger.With(zap.String("tg.callback_query", update.CallbackQuery.Data)))
	}
}

func LogIncoming(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, upd tgbotapi.Update, logger listener.ILogger) {
		logger.Info("incoming")
		next(ctx, upd, logger)
	}
}

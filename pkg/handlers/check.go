package handlers

import (
	"context"
	"strings"

	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CheckIfMessageNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.Message.Text)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

func CheckIfInlineQueryNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.InlineQuery.Query)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

func CheckIfCallbackQueryNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.CallbackQuery.Data)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

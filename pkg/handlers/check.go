package handlers

import (
	"context"
	"strings"

	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CheckIfMessageNotEmpty checks that incoming message is not empty, and stops handling if it's empty.
func CheckIfMessageNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.Message.Text)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

// CheckIfInlineQueryNotEmpty checks that incoming inline query is not empty, and stops handling if it's empty.
func CheckIfInlineQueryNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.InlineQuery.Query)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

// CheckIfCallbackQueryNotEmpty checks that incoming callback query is not empty, and stops handling if it's empty.
func CheckIfCallbackQueryNotEmpty(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, update tgbotapi.Update, logger listener.ILogger) {
		if len(strings.TrimSpace(update.CallbackQuery.Data)) == 0 {
			return
		}

		next(ctx, update, logger)
	}
}

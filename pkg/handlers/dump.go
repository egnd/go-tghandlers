package handlers

import (
	"context"

	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// DumpUpdate logs whole struct of incoming update.
func DumpUpdate(next listener.EventHandler) listener.EventHandler {
	return func(ctx context.Context, upd tgbotapi.Update, logger listener.ILogger) {
		logger.Debug("event", zap.Any("update", upd))
		next(ctx, upd, logger)
	}
}

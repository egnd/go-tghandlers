// Package listener ...
package listener

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// ILogger is interface for logger instance.
type ILogger interface {
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Debug(string, ...zap.Field)
	With(...zap.Field) *zap.Logger
}

// ITgAPI is interface for Telegram API instance.
type ITgAPI interface {
	GetUpdatesChan(tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	GetMe() (tgbotapi.User, error)
	GetChatMember(tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)
	LeaveChat(tgbotapi.ChatConfig) (tgbotapi.APIResponse, error)
	AnswerInlineQuery(tgbotapi.InlineConfig) (tgbotapi.APIResponse, error)
}

// EventHandler is telegram event handler callback interface.
type EventHandler func(context.Context, tgbotapi.Update, ILogger)

// EventHandlerDecorator is decorator interface for EventHandler, which provides handlers chaining ability.
type EventHandlerDecorator func(EventHandler) EventHandler

// ChainWrapper is callback interface for building chain of handlers from slice of decorators.
type ChainWrapper func([]EventHandlerDecorator, ChainWrapper) EventHandler

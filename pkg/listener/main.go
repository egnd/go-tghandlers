// Package listener ...
package listener

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

type ILogger interface {
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	Debug(string, ...zap.Field)
	With(...zap.Field) *zap.Logger
}

type ITgAPI interface {
	GetUpdatesChan(tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	GetMe() (tgbotapi.User, error)
	GetChatMember(tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)
	LeaveChat(tgbotapi.ChatConfig) (tgbotapi.APIResponse, error)
	AnswerInlineQuery(tgbotapi.InlineConfig) (tgbotapi.APIResponse, error)
}

type EventHandler func(context.Context, tgbotapi.Update, ILogger)

type EventHandlerDecorator func(EventHandler) EventHandler

type ChainWrapper func([]EventHandlerDecorator, ChainWrapper) EventHandler

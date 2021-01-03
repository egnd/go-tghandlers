package handlers_test

import (
	"context"
	"testing"

	"github.com/egnd/go-tghandlers/gen/mocks"
	"github.com/egnd/go-tghandlers/pkg/handlers"
	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func Test_AppendLogWithMessage(t *testing.T) {
	logger := &mocks.ILogger{}
	logger.On("With", mock.Anything).Return(zap.NewNop())
	var chained bool
	handlers.AppendLogWithMessage(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
		chained = true
	})(context.TODO(), tgbotapi.Update{Message: &tgbotapi.Message{}}, logger)
	assert.True(t, chained)
	logger.AssertExpectations(t)
}

func Test_AppendLogWithInlineQuery(t *testing.T) {
	logger := &mocks.ILogger{}
	logger.On("With", mock.Anything).Return(zap.NewNop())
	var chained bool
	handlers.AppendLogWithInlineQuery(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
		chained = true
	})(context.TODO(), tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{}}, logger)
	assert.True(t, chained)
	logger.AssertExpectations(t)
}

func Test_AppendLogWithCallbackQuery(t *testing.T) {
	logger := &mocks.ILogger{}
	logger.On("With", mock.Anything).Return(zap.NewNop())
	var chained bool
	handlers.AppendLogWithCallbackQuery(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
		chained = true
	})(context.TODO(), tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}}, logger)
	assert.True(t, chained)
	logger.AssertExpectations(t)
}

func Test_LogIncoming(t *testing.T) {
	logger := &mocks.ILogger{}
	logger.On("Info", mock.Anything).Return(zap.NewNop())
	var chained bool
	handlers.LogIncoming(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
		chained = true
	})(context.TODO(), tgbotapi.Update{}, logger)
	assert.True(t, chained)
	logger.AssertExpectations(t)
}

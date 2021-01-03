// Package handlers ...
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
)

func Test_DumpUpdate(t *testing.T) {
	logger := &mocks.ILogger{}
	logger.On("Debug", mock.Anything, mock.Anything)
	var chained bool
	handlers.DumpUpdate(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
		chained = true
	})(context.TODO(), tgbotapi.Update{}, logger)
	assert.True(t, chained)
	logger.AssertExpectations(t)
}

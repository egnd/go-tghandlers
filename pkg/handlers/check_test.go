package handlers_test

import (
	"context"
	"testing"

	"github.com/egnd/go-tghandlers/pkg/handlers"
	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_CheckIfMessageNotEmpty(t *testing.T) {
	cases := []struct {
		name    string
		chained bool
		upd     tgbotapi.Update
	}{
		{
			name:    "empty",
			upd:     tgbotapi.Update{Message: &tgbotapi.Message{}},
			chained: false,
		},
		{
			name:    "not empty",
			upd:     tgbotapi.Update{Message: &tgbotapi.Message{Text: "123"}},
			chained: true,
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			var chained bool
			handlers.CheckIfMessageNotEmpty(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
				chained = true
			})(context.TODO(), tcase.upd, zap.NewNop())
			assert.EqualValues(tt, tcase.chained, chained)
		})
	}
}

func Test_CheckIfInlineQueryNotEmpty(t *testing.T) {
	cases := []struct {
		name    string
		chained bool
		upd     tgbotapi.Update
	}{
		{
			name:    "empty",
			upd:     tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{}},
			chained: false,
		},
		{
			name:    "not empty",
			upd:     tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{Query: "123"}},
			chained: true,
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			var chained bool
			handlers.CheckIfInlineQueryNotEmpty(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
				chained = true
			})(context.TODO(), tcase.upd, zap.NewNop())
			assert.EqualValues(tt, tcase.chained, chained)
		})
	}
}

func Test_CheckIfCallbackQueryNotEmpty(t *testing.T) {
	cases := []struct {
		name    string
		chained bool
		upd     tgbotapi.Update
	}{
		{
			name:    "empty",
			upd:     tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}},
			chained: false,
		},
		{
			name:    "not empty",
			upd:     tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{Data: "123"}},
			chained: true,
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			var chained bool
			handlers.CheckIfCallbackQueryNotEmpty(func(ctx context.Context, upd tgbotapi.Update, loggr listener.ILogger) {
				chained = true
			})(context.TODO(), tcase.upd, zap.NewNop())
			assert.EqualValues(tt, tcase.chained, chained)
		})
	}
}

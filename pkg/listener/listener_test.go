package listener_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/egnd/go-tghandlers/gen/mocks"
	"github.com/egnd/go-tghandlers/pkg/listener"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_Listener(t *testing.T) {
	cases := []struct {
		name     string
		event    listener.EventType
		upd      tgbotapi.Update
		handled  bool
		err      error
		getMeErr error
	}{
		{
			name:    "EventDirectMessage",
			event:   listener.EventDirectMessage,
			upd:     tgbotapi.Update{Message: &tgbotapi.Message{}},
			handled: true,
		},
		{
			name:    "EventInlineQuery",
			event:   listener.EventInlineQuery,
			upd:     tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{}},
			handled: true,
		},
		{
			name:    "EventCallbackQuery",
			event:   listener.EventCallbackQuery,
			upd:     tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}},
			handled: true,
		},
		{
			name:    "EventUndefined",
			event:   listener.EventUndefined,
			handled: true,
		},
		{
			name:    "unexpected event",
			event:   listener.EventCallbackQuery,
			upd:     tgbotapi.Update{InlineQuery: &tgbotapi.InlineQuery{}},
			handled: false,
		},
		{
			name:     "GetMe error",
			event:    listener.EventDirectMessage,
			getMeErr: errors.New("error"),
			err:      errors.New("error"),
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			var handled bool
			api := &mocks.ITgAPI{}
			api.On("GetMe").Return(tgbotapi.User{}, tcase.getMeErr)
			obj := listener.NewListener()
			obj.Add(tcase.event, func(handler listener.EventHandler) listener.EventHandler {
				return func(ctx context.Context, upd tgbotapi.Update, lggr listener.ILogger) {
					handled = true
				}
			})
			updChan := make(chan tgbotapi.Update)
			var err error
			go func() {
				err = obj.Listen(context.TODO(), updChan, api, zap.NewNop())
			}()
			time.Sleep(time.Second)
			go func() {
				updChan <- tcase.upd
			}()
			time.Sleep(100 * time.Millisecond)
			assert.EqualValues(tt, tcase.handled, handled)
			assert.EqualValues(tt, tcase.err, err)
			api.AssertExpectations(tt)
		})
	}
}

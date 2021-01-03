package listener_test

import (
	"testing"

	"github.com/egnd/go-tghandlers/pkg/listener"
	"github.com/stretchr/testify/assert"
)

func Test_GetEventName(t *testing.T) {
	cases := []struct {
		name   string
		typeID listener.EventType
	}{
		{"direct_message", listener.EventDirectMessage},
		{"inline_query", listener.EventInlineQuery},
		{"callback_query", listener.EventCallbackQuery},
		{name: "undefined"},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			assert.EqualValues(tt, tcase.name, listener.GetEventName(tcase.typeID))
		})
	}
}

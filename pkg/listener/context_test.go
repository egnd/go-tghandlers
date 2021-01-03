package listener_test

import (
	"context"
	"testing"

	"github.com/egnd/go-tghandlers/pkg/listener"
	"github.com/stretchr/testify/assert"
)

func Test_GetEventTypeFromCtx(t *testing.T) {
	assert.EqualValues(t, listener.EventInlineQuery, listener.GetEventTypeFromCtx(
		context.WithValue(context.TODO(), listener.CtxEventType, listener.EventInlineQuery),
	))
}

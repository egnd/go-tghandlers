package listener

import (
	"context"
)

type ctxKeyType int

const (
	// CtxEventType context key for storing event type ID.
	CtxEventType ctxKeyType = iota
)

// GetEventTypeFromCtx return event ID from context struct.
func GetEventTypeFromCtx(ctx context.Context) EventType {
	return ctx.Value(CtxEventType).(EventType)
}

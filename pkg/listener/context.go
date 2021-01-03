package listener

import (
	"context"
)

type ctxKeyType int

const (
	ctxEventType ctxKeyType = iota
)

func GetEventTypeFromCtx(ctx context.Context) EventType {
	return ctx.Value(ctxEventType).(EventType)
}

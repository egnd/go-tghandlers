package listener

import (
	"context"
)

type ctxKeyType int

const (
	CtxEventType ctxKeyType = iota
)

func GetEventTypeFromCtx(ctx context.Context) EventType {
	return ctx.Value(CtxEventType).(EventType)
}

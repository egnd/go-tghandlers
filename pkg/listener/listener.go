package listener

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/zap"
)

// Listener is struct which is listening and handling events.
type Listener struct {
	decorators map[EventType][]EventHandlerDecorator
	handlers   map[EventType]EventHandler
}

// NewListener constructor for Listener struct.
func NewListener() *Listener {
	return &Listener{
		decorators: make(map[EventType][]EventHandlerDecorator),
	}
}

// Add adds decorators for handling specific Telegram event.
func (b *Listener) Add(eType EventType, decorators ...EventHandlerDecorator) {
	if _, ok := b.decorators[eType]; !ok {
		b.decorators[eType] = []EventHandlerDecorator{}
	}

	b.decorators[eType] = append(b.decorators[eType], decorators...)
}

func (b *Listener) wrapDecoratorsChain(
	decorators []EventHandlerDecorator,
	wrapper ChainWrapper,
) (handler EventHandler) {
	if len(decorators) == 0 {
		return func(ctx context.Context, upd tgbotapi.Update, logger ILogger) {}
	}

	return decorators[0](wrapper(decorators[1:], wrapper))
}

func (b *Listener) buildHandlers() {
	b.handlers = make(map[EventType]EventHandler)

	for eType, decorators := range b.decorators {
		b.handlers[eType] = b.wrapDecoratorsChain(decorators, b.wrapDecoratorsChain)
	}
}

func (b *Listener) defineType(update tgbotapi.Update) (typeCode EventType) {
	switch {
	case update.Message != nil:
		typeCode = EventDirectMessage
	case update.InlineQuery != nil:
		typeCode = EventInlineQuery
	case update.CallbackQuery != nil:
		typeCode = EventCallbackQuery
	default:
		typeCode = EventUndefined
	}

	return
}

// Listen starts listening incoming messages in channel.
func (b *Listener) Listen(ctx context.Context,
	updChan tgbotapi.UpdatesChannel,
	api ITgAPI,
	logger ILogger,
) (err error) {
	b.buildHandlers()

	var me tgbotapi.User

	if me, err = api.GetMe(); err != nil {
		return
	}

	logger.Info("listening...", zap.Any("bot", me))

	for update := range updChan { // @TODO: add workers pool
		eType := b.defineType(update)
		eventLogger := logger.With(zap.String("tg.event", GetEventName(eType)))

		handler, ok := b.handlers[eType]
		if !ok {
			eventLogger.Warn("unexpected event", zap.Any("data", update))

			continue
		}

		handler(context.WithValue(ctx, CtxEventType, eType), update, eventLogger)
	}

	return
}

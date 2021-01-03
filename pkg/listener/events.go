package listener

// EventType variable type for event type ID.
type EventType int

const (
	// EventUndefined undefined event.
	EventUndefined EventType = iota
	// EventDirectMessage receiving direct message event.
	EventDirectMessage
	// EventInlineQuery receiving inline query event.
	EventInlineQuery
	// EventCallbackQuery receiving callback query event.
	EventCallbackQuery
)

// GetEventName return event name by it's ID.
func GetEventName(eventID EventType) string {
	switch eventID {
	case EventDirectMessage:
		return "direct_message"
	case EventInlineQuery:
		return "inline_query"
	case EventCallbackQuery:
		return "callback_query"
	case EventUndefined:
		fallthrough
	default:
		return "undefined"
	}
}

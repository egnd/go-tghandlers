package listener

type EventType int

const (
	EventUndefined EventType = iota
	EventDirectMessage
	EventInlineQuery
	EventCallbackQuery
)

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

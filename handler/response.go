package handler

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
}

func newResponse(messageType, message string, data any) response {
	return response{
		messageType,
		message,
		data,
	}
}

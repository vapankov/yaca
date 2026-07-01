package http

// General codes and messages.

const (
	codeRequestHasNoBody    = 1
	codeInternalServerError = 2
)

const (
	messageRequestHasNoBody    = "Request has no body."
	messageInternalServerError = "Internal server error."
)

// Post message codes and messages.

const (
	codeEmptyMessage = 100
)

const (
	messageEmptyMessageContents = "Message contents is empty."
)

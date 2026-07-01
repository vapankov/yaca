package http

import (
	"context"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
)

type Usecases interface {
	PostMessage(ctx context.Context, params *usecases.PostMessageInput) (*usecases.PostMessageOutput, error)
	ViewMessages(ctx context.Context, params *usecases.ViewMessagesInput) (*usecases.ViewMessagesOutput, error)
}

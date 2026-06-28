package usecases

import (
	"context"
	"time"

	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type MessageRespository interface {
	CreateMessage(ctx context.Context, message *entities.Message) error
	SearchMessages(ctx context.Context, params *SearchMessagesQuery) ([]*entities.Message, error)
}

type MessageIDGenerator interface {
	Generate() values.MessageID
}

type Clock interface {
	Now() time.Time
}

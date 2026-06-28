package indentifiers

import (
	"github.com/google/uuid"

	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type Message struct {
}

func NewMessage() Message {
	return Message{}
}

func (Message) Generate() values.MessageID {
	return values.MessageID(uuid.New().String())
}

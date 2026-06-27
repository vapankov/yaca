package entities

import (
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type Message struct {
	ID values.MessageID

	Contents values.MessageContents
	Metadata *values.MessageMetadata
}

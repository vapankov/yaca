package values

import "time"

type (
	MessageID        string
	MessageContents  string
	MessageCreatedAt time.Time

	MessageMetadata struct {
		CreatedAt MessageCreatedAt
	}
)

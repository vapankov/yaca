package usecases

import (
	"context"
	"fmt"

	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type PostMessageParams struct {
	MessageContents values.MessageContents
}

type PostMessageResult struct {
}

func (ucs *UseCases) PostMessage(ctx context.Context, params *PostMessageParams) (*PostMessageResult, error) {
	var (
		messageID = ucs.messageIDGenerator.GenerateMessageID()
		contents  = params.MessageContents
		createdAt = values.MessageCreatedAt(ucs.timeProvider.TimeNow())

		message = entities.Message{
			ID:       messageID,
			Contents: contents,
			Metadata: &values.MessageMetadata{
				CreatedAt: createdAt,
			},
		}
	)

	if err := ucs.messageRepository.CreateMessage(ctx, &message); err != nil {
		return nil, fmt.Errorf("repository: create message: %w", err)
	}

	return &PostMessageResult{}, nil
}

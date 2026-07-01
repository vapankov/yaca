package usecases

import (
	"context"
	"fmt"

	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type PostMessageInput struct {
	MessageContents values.MessageContents
}

type PostMessageOutput struct {
}

func (ucs *UseCases) PostMessage(ctx context.Context, input *PostMessageInput) (*PostMessageOutput, error) {
	var (
		messageID = ucs.messageIDGenerator.Generate()
		contents  = input.MessageContents
		createdAt = values.MessageCreatedAt(ucs.clock.Now())

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

	return &PostMessageOutput{}, nil
}

package usecases

import (
	"context"
	"fmt"

	"github.com/vapankov/yaca/internal/types"

	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type ViewMessagesInput struct {
	Pagination *types.Pagination
}

type (
	ViewMessagesOutput struct {
		Items []ViewMessagesResultItem
	}

	ViewMessagesResultItem struct {
		MessageID        values.MessageID
		MessageContents  values.MessageContents
		MessageCreatedAt values.MessageCreatedAt
	}
)

func (ucs *UseCases) ViewMessages(ctx context.Context, input *ViewMessagesInput) (*ViewMessagesOutput, error) {
	messages, err := ucs.messageRepository.SearchMessages(ctx, &SearchMessagesQuery{
		Sort: &SearchMessagesQuerySort{
			CreatedAt: types.OrderDesc,
		},
		Pagination: input.Pagination,
	})
	if err != nil {
		return nil, fmt.Errorf("repository: search messages: %w", err)
	}

	items := make([]ViewMessagesResultItem, len(messages))
	for i := range messages {
		items[i] = ViewMessagesResultItem{
			MessageID:        messages[i].ID,
			MessageContents:  messages[i].Contents,
			MessageCreatedAt: messages[i].Metadata.CreatedAt,
		}
	}

	return &ViewMessagesOutput{
		Items: items,
	}, nil
}

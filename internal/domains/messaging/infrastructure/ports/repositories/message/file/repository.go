package file

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/libs/pagination"
	"github.com/vapankov/yaca/internal/types"
)

type Repository struct {
	lineStorage FileLineStorage
}

func New(lineStorage FileLineStorage) *Repository {
	return &Repository{
		lineStorage: lineStorage,
	}
}

func (r *Repository) CreateMessage(ctx context.Context, message *entities.Message) error {
	if message == nil {
		return nil
	}

	dto := encodeMessage(message)

	line, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	if err := r.lineStorage.Insert(string(line)); err != nil {
		return fmt.Errorf("insert line: %w", err)
	}

	return nil
}

func (r *Repository) SearchMessages(ctx context.Context, params *usecases.SearchMessagesQuery) ([]*entities.Message, error) {
	lines, err := r.lineStorage.Read()
	if err != nil {
		return nil, fmt.Errorf("read lines: %w", err)
	}

	var messages []*entities.Message

	for _, line := range lines {
		var dto messageDTO
		if err := json.Unmarshal([]byte(line), &dto); err != nil {
			return nil, fmt.Errorf("json unmarshal: %w", err)
		}

		message, err := decodeMessage(&dto)
		if err != nil {
			return nil, fmt.Errorf("decode message: %w", err)
		}

		messages = append(messages, message)
	}

	if params != nil && params.Sort != nil {
		switch params.Sort.CreatedAt {
		case types.OrderAsc:
			slices.SortFunc(messages, func(a, b *entities.Message) int {
				return compareMessagesByCreatedAt(a, b)
			})
		case types.OrderDesc:
			slices.SortFunc(messages, func(a, b *entities.Message) int {
				return compareMessagesByCreatedAt(b, a)
			})
		}
	}

	if params != nil && params.Pagination != nil {
		messages = pagination.PaginateSlice(messages, params.Pagination)
	}

	return messages, nil
}

func compareMessagesByCreatedAt(a, b *entities.Message) int {
	if a == nil && b == nil {
		return 0
	}

	if a == nil {
		return -1
	}

	if b == nil {
		return 1
	}

	if a.Metadata == nil && b.Metadata == nil {
		return 0
	}

	if a.Metadata == nil {
		return -1
	}

	if b.Metadata == nil {
		return 1
	}

	return int(time.Time(a.Metadata.CreatedAt).Sub(time.Time(b.Metadata.CreatedAt)))
}

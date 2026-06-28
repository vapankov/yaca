package file

import (
	"fmt"
	"time"

	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

type messageDTO struct {
	ID       string              `json:"id"`
	Contents string              `json:"contents"`
	Metadata *messageMetadataDTO `json:"metadata"`
}

func encodeMessage(entity *entities.Message) *messageDTO {
	if entity == nil {
		return nil
	}

	return &messageDTO{
		ID:       string(entity.ID),
		Contents: string(entity.Contents),
		Metadata: encodeMessageMetadata(entity.Metadata),
	}
}

func decodeMessage(dto *messageDTO) (*entities.Message, error) {
	if dto == nil {
		return nil, nil
	}

	metadata, err := decodeMessageMetadata(dto.Metadata)
	if err != nil {
		return nil, fmt.Errorf("decode metadata: %w", err)
	}

	return &entities.Message{
		ID:       values.MessageID(dto.ID),
		Contents: values.MessageContents(dto.Contents),
		Metadata: metadata,
	}, nil
}

type messageMetadataDTO struct {
	CreatedAt string `json:"created_at"`
}

func encodeMessageMetadata(val *values.MessageMetadata) *messageMetadataDTO {
	if val == nil {
		return nil
	}

	return &messageMetadataDTO{
		CreatedAt: time.Time(val.CreatedAt).Format(time.RFC3339),
	}
}

func decodeMessageMetadata(dto *messageMetadataDTO) (*values.MessageMetadata, error) {
	if dto == nil {
		return nil, nil
	}

	createdAt, err := time.Parse(time.RFC3339, dto.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("parse created_at: %w", err)
	}

	return &values.MessageMetadata{
		CreatedAt: values.MessageCreatedAt(createdAt),
	}, nil
}

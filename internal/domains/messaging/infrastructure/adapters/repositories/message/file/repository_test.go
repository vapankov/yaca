package file

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
	"github.com/vapankov/yaca/internal/domains/messaging/core/entities"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
	"github.com/vapankov/yaca/internal/types"
)

type repoMocks struct {
	lineStorage *mockFileLineStorage
}

func newMockedRepo(t *testing.T) (*Repository, *repoMocks) {
	var (
		lineStorage = newMockFileLineStorage(t)
	)

	return New(lineStorage), &repoMocks{
		lineStorage: lineStorage,
	}
}

func TestRepository(t *testing.T) {
	t.Parallel()

	mkTime := func(dateTimeString string) time.Time {
		parsedTime, err := time.Parse(time.RFC3339, dateTimeString)
		require.NoError(t, err)

		return parsedTime
	}

	mkCreatedAt := func(dateTimeString string) values.MessageCreatedAt {
		return values.MessageCreatedAt(mkTime(dateTimeString))
	}

	testErr := errors.New("test error")

	t.Run("storage error on create", func(t *testing.T) {
		t.Parallel()

		var (
			message = &entities.Message{
				ID:       "123",
				Contents: "123",
			}

			messageDTO = &messageDTO{
				ID:       "123",
				Contents: "123",
			}
		)

		repo, mocks := newMockedRepo(t)

		messageLineBytes, err := json.Marshal(&messageDTO)
		require.NoError(t, err)

		mocks.lineStorage.EXPECT().Insert(string(messageLineBytes)).Return(testErr)

		err = repo.CreateMessage(t.Context(), message)
		require.ErrorIs(t, err, testErr)
	})

	t.Run("successful create", func(t *testing.T) {
		t.Parallel()

		var (
			message = &entities.Message{
				ID:       "123",
				Contents: "123",
			}

			messageDTO = &messageDTO{
				ID:       "123",
				Contents: "123",
			}
		)

		repo, mocks := newMockedRepo(t)

		messageLineBytes, err := json.Marshal(&messageDTO)
		require.NoError(t, err)

		mocks.lineStorage.EXPECT().Insert(string(messageLineBytes)).Return(nil)

		err = repo.CreateMessage(t.Context(), message)
		require.NoError(t, err)
	})

	t.Run("storage error on search", func(t *testing.T) {
		t.Parallel()

		repo, mocks := newMockedRepo(t)

		mocks.lineStorage.EXPECT().Read().Return(nil, testErr)

		actualMessages, err := repo.SearchMessages(t.Context(), nil)
		require.ErrorIs(t, err, testErr)
		require.Empty(t, actualMessages)
	})

	t.Run("empty repository", func(t *testing.T) {
		t.Parallel()

		repo, mocks := newMockedRepo(t)

		mocks.lineStorage.EXPECT().Read().Return(nil, nil)

		actualMessages, err := repo.SearchMessages(t.Context(), nil)
		require.NoError(t, err)
		require.Empty(t, actualMessages)
	})

	t.Run("searching", func(t *testing.T) {
		t.Parallel()

		var (
			messages = []*entities.Message{
				{
					ID:       "1",
					Contents: "1",
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &values.MessageMetadata{
						CreatedAt: mkCreatedAt("2026-06-28T00:00:00Z"),
					},
				},
				{
					ID:       "3",
					Contents: "3",
					Metadata: &values.MessageMetadata{},
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &values.MessageMetadata{
						CreatedAt: mkCreatedAt("2026-06-28T02:00:00Z"),
					},
				},
				{
					ID:       "4",
					Contents: "4",
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &values.MessageMetadata{
						CreatedAt: mkCreatedAt("2026-06-28T01:00:00Z"),
					},
				},
			}

			messageDTOs = []*messageDTO{
				{
					ID:       "1",
					Contents: "1",
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &messageMetadataDTO{
						CreatedAt: "2026-06-28T00:00:00Z",
					},
				},
				{
					ID:       "3",
					Contents: "3",
					Metadata: &messageMetadataDTO{
						CreatedAt: "0001-01-01T00:00:00Z",
					},
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &messageMetadataDTO{
						CreatedAt: "2026-06-28T02:00:00Z",
					},
				},
				{
					ID:       "4",
					Contents: "4",
				},
				{
					ID:       "2",
					Contents: "2",
					Metadata: &messageMetadataDTO{
						CreatedAt: "2026-06-28T01:00:00Z",
					},
				},
			}
		)

		messageDTOLines := make([]string, 0, len(messageDTOs))
		for _, dto := range messageDTOs {
			lineBytes, err := json.Marshal(dto)
			require.NoError(t, err)

			messageDTOLines = append(messageDTOLines, string(lineBytes))
		}

		t.Run("nil query", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), nil)

			require.NoError(t, err)
			require.Equal(t, messages, actualMessages)
		})

		t.Run("empty query", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{})

			require.NoError(t, err)
			require.Equal(t, messages, actualMessages)
		})

		t.Run("sort order is not available", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Sort: &usecases.SearchMessagesQuerySort{},
			})

			require.NoError(t, err)
			require.Equal(t, messages, actualMessages)
		})

		t.Run("sort by created_at asc", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Sort: &usecases.SearchMessagesQuerySort{
					CreatedAt: types.OrderAsc,
				},
			})

			expectedMessages := []*entities.Message{
				messages[0],
				messages[4],
				messages[2],
				messages[1],
				messages[5],
				messages[3],
			}

			require.NoError(t, err)
			require.Equal(t, expectedMessages, actualMessages)
		})

		t.Run("sort by created_at dsc", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Sort: &usecases.SearchMessagesQuerySort{
					CreatedAt: types.OrderDesc,
				},
			})

			expectedMessages := []*entities.Message{
				messages[3],
				messages[5],
				messages[1],
				messages[2],
				messages[0],
				messages[4],
			}

			require.NoError(t, err)
			require.Equal(t, expectedMessages, actualMessages)
		})

		t.Run("empty pagination", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Pagination: &types.Pagination{},
			})

			expectedMessages := []*entities.Message{}

			require.NoError(t, err)
			require.Equal(t, expectedMessages, actualMessages)
		})

		t.Run("sort by created_at asc and get second page", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Sort: &usecases.SearchMessagesQuerySort{
					CreatedAt: types.OrderAsc,
				},
				Pagination: &types.Pagination{
					PageNumber: 1,
					PageSize:   2,
				},
			})

			expectedMessages := []*entities.Message{
				messages[2],
				messages[1],
			}

			require.NoError(t, err)
			require.Equal(t, expectedMessages, actualMessages)
		})

		t.Run("sort by created_at dsc and get part of last page", func(t *testing.T) {
			t.Parallel()

			repo, mocks := newMockedRepo(t)

			mocks.lineStorage.EXPECT().Read().Return(messageDTOLines, nil)

			actualMessages, err := repo.SearchMessages(t.Context(), &usecases.SearchMessagesQuery{
				Sort: &usecases.SearchMessagesQuerySort{
					CreatedAt: types.OrderDesc,
				},
				Pagination: &types.Pagination{
					PageNumber: 1,
					PageSize:   4,
				},
			})

			expectedMessages := []*entities.Message{
				messages[0],
				messages[4],
			}

			require.NoError(t, err)
			require.Equal(t, expectedMessages, actualMessages)
		})
	})
}

package usecases

import "github.com/vapankov/yaca/internal/types"

type (
	SearchMessagesQuery struct {
		Filter     *SearchMessagesQueryFilter
		Sort       *SearchMessagesQuerySort
		Pagination *types.Pagination
	}

	SearchMessagesQueryFilter struct {
	}

	SearchMessagesQuerySort struct {
		CreatedAt types.Order
	}
)

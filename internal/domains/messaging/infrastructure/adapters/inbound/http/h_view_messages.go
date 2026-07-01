package http

import (
	"context"
	"net/http"
	"time"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
	"github.com/vapankov/yaca/internal/types"
)

const (
	viewMessagesDefaultPageNumber uint = 1
	viewMessagesDefaultPageSize   uint = 10
)

func (hs *Handlers) ViewMessages(ctx context.Context, request ViewMessagesRequestObject) (ViewMessagesResponseObject, error) {
	pageNumber := viewMessagesDefaultPageNumber
	if request.Params.PaginationPageNumber != nil {
		pageNumber = uint(*request.Params.PaginationPageNumber)
	}

	pageSize := viewMessagesDefaultPageSize
	if request.Params.PaginationPageSize != nil {
		pageSize = uint(*request.Params.PaginationPageSize)
	}

	input := &usecases.ViewMessagesInput{
		Pagination: &types.Pagination{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}

	output, err := hs.usecases.ViewMessages(ctx, input)
	if err != nil {
		return ViewMessagesdefaultJSONResponse{
			Body: ErrorResponse{
				Code:    codeInternalServerError,
				Message: messageInternalServerError,
			},
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	items := make([]ViewMessagesResponseItem, len(output.Items))
	for i, item := range output.Items {
		items[i] = ViewMessagesResponseItem{
			Id:        string(item.MessageID),
			Contents:  string(item.MessageContents),
			CreatedAt: time.Time(item.MessageCreatedAt),
		}
	}

	return ViewMessages200JSONResponse{
		Items: items,
	}, nil
}

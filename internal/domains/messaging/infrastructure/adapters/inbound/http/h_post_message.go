package http

import (
	"context"
	"net/http"

	"github.com/vapankov/yaca/internal/domains/messaging/application/usecases"
	"github.com/vapankov/yaca/internal/domains/messaging/core/values"
)

func (hs *Handlers) PostMessage(ctx context.Context, request PostMessageRequestObject) (PostMessageResponseObject, error) {
	if request.Body == nil {
		return PostMessagedefaultJSONResponse{
			Body: ErrorResponse{
				Code:    codeRequestHasNoBody,
				Message: messageRequestHasNoBody,
			},
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	if request.Body.Contents == "" {
		return PostMessagedefaultJSONResponse{
			Body: ErrorResponse{
				Code:    codeEmptyMessage,
				Message: messageEmptyMessageContents,
			},
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	input := &usecases.PostMessageInput{
		MessageContents: values.MessageContents(request.Body.Contents),
	}

	_, err := hs.usecases.PostMessage(ctx, input)
	if err != nil {
		return PostMessagedefaultJSONResponse{
			Body: ErrorResponse{
				Code:    codeInternalServerError,
				Message: messageInternalServerError,
			},
		}, nil
	}

	return PostMessage200JSONResponse{}, nil
}

package usecases

type UseCases struct {
	messageRepository  MessageRespository
	messageIDGenerator MessageIDGenerator
	timeProvider       TimeProvider
}

func New(
	messageRepository MessageRespository,
	messageIDGenerator MessageIDGenerator,
	timeProvider TimeProvider,
) *UseCases {
	return &UseCases{
		messageRepository:  messageRepository,
		messageIDGenerator: messageIDGenerator,
		timeProvider:       timeProvider,
	}
}

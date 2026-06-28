package usecases

type UseCases struct {
	messageRepository  MessageRespository
	messageIDGenerator MessageIDGenerator
	clock              Clock
}

func New(
	messageRepository MessageRespository,
	messageIDGenerator MessageIDGenerator,
	clock Clock,
) *UseCases {
	return &UseCases{
		messageRepository:  messageRepository,
		messageIDGenerator: messageIDGenerator,
		clock:              clock,
	}
}

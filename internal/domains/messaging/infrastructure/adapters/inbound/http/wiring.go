package http

type Handlers struct {
	usecases Usecases
}

func NewHandlers(usecases Usecases) *Handlers {
	return &Handlers{usecases: usecases}
}

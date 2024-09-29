package interfaces

import (
	"fmt"
	"net/http"

	"github.com/felonardo/leosoftware/usecase"
)

// GreetingHandler handles HTTP requests for greetings
type GreetingHandler struct {
	UseCase *usecase.GreetingUseCase
}

// NewGreetingHandler creates a new GreetingHandler
func NewGreetingHandler(useCase *usecase.GreetingUseCase) *GreetingHandler {
	return &GreetingHandler{UseCase: useCase}
}

// ServeHTTP handles the greeting request and sends the response
func (h *GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	greeting := h.UseCase.GetGreeting()
	fmt.Fprintf(w, greeting)
}

// usecase/greeting_usecase.go
package usecase

import (
	"github.com/felonardo/leosoftware/model"
)

// GreetingRepository defines the interface for data access
type GreetingRepository interface {
	GetGreeting() model.Greeting
}

// GreetingUseCase orchestrates greeting logic
type GreetingUseCase struct {
	repo GreetingRepository
}

// NewGreetingUseCase creates a new instance of GreetingUseCase
func NewGreetingUseCase(repo GreetingRepository) *GreetingUseCase {
	return &GreetingUseCase{repo: repo}
}

// GetGreeting returns the greeting message
func (uc *GreetingUseCase) GetGreeting() string {
	greeting := uc.repo.GetGreeting()
	return greeting.Message
}

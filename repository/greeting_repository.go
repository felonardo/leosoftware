// repository/greeting_repository.go
package repository

import "github.com/felonardo/leosoftware/model"

// InMemoryGreetingRepository is an in-memory implementation of GreetingRepository
type InMemoryGreetingRepository struct{}

// NewInMemoryGreetingRepository creates a new instance of InMemoryGreetingRepository
func NewInMemoryGreetingRepository() *InMemoryGreetingRepository {
	return &InMemoryGreetingRepository{}
}

// GetGreeting returns a static greeting
func (repo *InMemoryGreetingRepository) GetGreeting() model.Greeting {
	return model.Greeting{Message: "Hello, World!"}
}

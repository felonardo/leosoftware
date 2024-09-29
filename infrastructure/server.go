// infrastructure/server.go
package infrastructure

import (
	"log"
	"net/http"

	"github.com/felonardo/leosoftware/interfaces"
	"github.com/felonardo/leosoftware/repository"
	"github.com/felonardo/leosoftware/usecase"
)

// StartServer initializes and starts the HTTP server
func StartServer() {
	repo := repository.NewInMemoryGreetingRepository()
	uc := usecase.NewGreetingUseCase(repo)
	handler := interfaces.NewGreetingHandler(uc)

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	log.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}

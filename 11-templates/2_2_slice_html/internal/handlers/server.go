package handlers

import (
	"context"
	"net/http"
	"time"

	"learngo-pockets/templates/internal/habit"
)

// Client is the dependency towards the Habits service.
//
//go:generate minimock -i habitsClient -s "_mock.go" -o "mocks"
type habitsClient interface {
	ListHabits(ctx context.Context, t time.Time) ([]habit.Habit, error)
}

// Server serves all the HTML routes on this service.
type Server struct {
	client habitsClient
	lgr    Logger
}

// New builds a new server.
func New(cli habitsClient, lgr Logger) *Server {
	return &Server{
		client: cli,
		lgr:    lgr,
	}
}

// Router returns an HTTP handler that listens to all the proper paths.
func (s *Server) Router() http.Handler {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodGet+" "+"/", s.index)

	return r
}

type Logger interface {
	Logf(format string, args ...any)
}

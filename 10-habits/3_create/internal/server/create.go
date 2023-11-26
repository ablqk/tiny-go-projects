package server

import (
	"context"
	"errors"
	"fmt"
	"log"

	"learngo-pockets/habits/api"
	"learngo-pockets/habits/internal/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateHabit is the endpoint that registers a habit.
func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	log.Printf("CreateHabit request received: %s", request)

	err := validateCreateHabitRequest(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request: "+err.Error())
	}

	var freq uint
	if request.WeeklyFrequency != nil && uint(*request.WeeklyFrequency) > 0 {
		freq = uint(*request.WeeklyFrequency)
	}

	h := habit.Habit{
		Name:            habit.Name(request.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}

	savedHabit, err := habit.CreateHabit(ctx, h)
	if err != nil {
		invalidErr := habit.InvalidInputError{}
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		// other error
		return nil, fmt.Errorf("cannot save habit %v: %w", h, err)
	}

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id:              string(savedHabit.ID),
			Name:            string(savedHabit.Name),
			WeeklyFrequency: int32(savedHabit.WeeklyFrequency),
		},
	}, nil
}

func validateCreateHabitRequest(request *api.CreateHabitRequest) error {
	switch {
	case request == nil:
		return fmt.Errorf("empty request")
	case request.Name == "":
		return fmt.Errorf("missing name of habit")
	}
	return nil
}
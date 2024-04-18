package main

import (
	"context"
	"log"

	"learngo-pockets/habits/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial("localhost:8084", creds)
	if err != nil {
		log.Fatal(err)
	}

	mock(api.NewHabitsClient(conn))
}

func mock(cli api.HabitsClient) {
	cli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            "Call your mom",
		WeeklyFrequency: ptr(1),
	})
	cli.CreateHabit(context.Background(), &api.CreateHabitRequest{
		Name:            "Sleep",
		WeeklyFrequency: ptr(7),
	})
}

func ptr(i int32) *int32 {
	return &i
}
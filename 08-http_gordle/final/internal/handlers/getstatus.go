package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"learngo-pockets/httpgordle/api"
)

// getStatus is the HTTP implementation of the endpoint that returns the status of a game.
func getStatus(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params[api.GameID]
	log.Printf("retrieve status from id: %v", id)

	// TODO: retrieve status from game id
	response := api.GameResponse{
		ID:           "1682279480",
		AttemptsLeft: 0,
		Guesses: []api.Guess{
			{"hello", "⬜️⬜️🟡⬜️🟡"},
			{"sauna", "⬜️⬜️🟡⬜️⬜️"},
			{"cloud", "💚💚💚💚💚"}},
		Solution: "cloud",
	}

	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "failed to write response", http.StatusInternalServerError)
	}
}

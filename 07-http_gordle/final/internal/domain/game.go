package domain

import "learngo-pockets/httpgordle/internal/gordle"

// A GameID represents an identified of a game
type GameID string

// Game contains the information about a game.
type Game struct {
	// ID is the identified of a game.
	ID     GameID
	Gordle gordle.Game

	// AttemptsLeft counts the number of attempts left before the game is over.
	AttemptsLeft byte
	// Guesses is the list of past guesses, and their feedback.
	Guesses []Guess
	// Solution is Gordle's secret word. It is only provided when there are no attempts left.
	Solution string
}

// A Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string
	Feedback string
}

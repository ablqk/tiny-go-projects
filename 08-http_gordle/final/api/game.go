package api

// GameResponse contains the information about a game.
type GameResponse struct {
	// ID is the identified of a game.
	ID string `json:"id"`
	// AttemptsLeft counts the number of attempts left before the game is over.
	AttemptsLeft byte `json:"attempts_left"`
	// Guesses is the list of past guesses, and their feedback.
	Guesses []Guess `json:"guesses"`
	// Solution is Gordle's secret word. It is only provided when there are no attempts left.
	Solution string `json:"solution,omitempty"`
}

// A Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}

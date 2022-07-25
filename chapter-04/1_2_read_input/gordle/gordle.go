package gordle

import (
	"fmt"
	"os"
)

const wordLength = 5

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
}

// New returns a Gordle variable, which can be used to Play!
func New() *Gordle {
	g := &Gordle{}

	return g
}

// Play runs the game.
func (g *Gordle) Play() {
	fmt.Printf("Enter a %d-letter guess:\n", wordLength)

	var (
		attempt        []rune
		attemptIsValid bool
	)

	for !attemptIsValid {
		// Read the attempt from the player.
		wordCount, err := fmt.Fscanf(os.Stdin, "%s", &attempt)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's word: %q", err)
			continue
		}
		// We expect a single word.
		if wordCount != 1 {
			fmt.Fprintf(os.Stderr, "error while reading the player's word: a single word wasn't provided, got %d instead", wordCount)
			continue
		}

		// Verify the suggestion has a valid length.
		if len(attempt) != wordLength {
			fmt.Errorf("invalid word length: expected %d, got %d", wordLength, len(attempt))
		} else {
			attemptIsValid = true
		}
	}

	fmt.Printf("Your guess: %q\n", attempt)
}

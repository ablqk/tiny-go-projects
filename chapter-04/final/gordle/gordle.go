package gordle

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// New returns a Gordle variable, which can be used to Play!
func New(cfs ...ConfigFunc) (*Gordle, error) {
	g := &Gordle{
		scanner:     bufio.NewScanner(os.Stdin), // read from stdin by default
		maxAttempts: -1,                         // no maximum number of attempts by default
		solution:    randomWord(),               // pick a random word from the corpus
	}

	for _, cf := range cfs {
		err := cf(g)
		if err != nil {
			return nil, fmt.Errorf("unable to apply config func: %w", err)
		}
	}

	// delay the checker creation till here, in case the solution was passed as a config func.
	g.solutionChecker = &solutionChecker{solution: g.solution}
	return g, nil
}

// Play runs the game. It will exit when the maximum number of attempts was reached, or if the word was found.
func (g *Gordle) Play() {
	// break condition: we've reached the maximum number of attempts
	for g.currentAttempt != g.maxAttempts {
		// ask for a valid word
		word := g.ask()

		// check it
		fb := g.solutionChecker.check(word)

		// print the feedback
		fmt.Println(fb)
		if string(word) == string(g.solution) {
			fmt.Printf("🎉 You won! You found in %d attempt(s)! The word was: %s.\n", g.currentAttempt, string(g.solution))
			return
		}
		g.currentAttempt++
	}
	// we've exhausted the number of allowed attempts
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// Gordle holds all the information we need to play a game of gordle.
type Gordle struct {
	scanner         scanner
	solution        []rune
	maxAttempts     int
	currentAttempt  int
	solutionChecker *solutionChecker
}

// ask scans until a valid suggestion is made (and returned).
func (g *Gordle) ask() []rune {
	fmt.Println("Enter a guess:")

	for g.scanner.Scan() {
		suggestion := g.scanner.Text()
		if g.scanner.Err() != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error while reading the player's word: %q", g.scanner.Err())
			continue
		}

		attempt := []rune(strings.ToUpper(suggestion))
		err := g.validateAttempt(attempt)
		if err != nil {
			fmt.Println(err)
		} else {
			return attempt
		}
	}
	// this can't happen
	return []rune{}
}

var errInvalidWordLength = fmt.Errorf("invalid attempt, word doesn't have the same number of letters as the solution ")

// validateAttempt ensures the attempt is valid enough.
func (g *Gordle) validateAttempt(attempt []rune) error {
	if len(attempt) != len(g.solution) {
		return errInvalidWordLength
	}
	return nil
}

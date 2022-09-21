package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/3_corpus/gordle"
)

const maxAttempts = 6

func main() {
	fmt.Println("Welcome to Gordle!")

	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g := gordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)

	// Run the game ! It will end when it's over.
	g.Play()
}

package main

import (
	"fmt"
	"log/slog"
	"os"

	"learngo/09/maze/internal/solver"
)

// main takes 2 parameters when run:
// first, the path to a PNG image that contains the maze.
//
//	The colours of walls, path, treasure, etc. can be configured in internal/solver/config.go.
//
// second, the path where it should write the solution in a similar PNG image.
func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	slog.Info(fmt.Sprintf("Solving maze %q and saving it as %q", inputFile, outputFile))

	s, err := solver.New(inputFile)
	if err != nil {
		exit(err)
	}

	err = s.Solve()
	if err != nil {
		exit(err)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		exit(err)
	}
}

// usage displays the usage of the binary and exits the program.
func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}

// exit prints the error and exits the program.
func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}

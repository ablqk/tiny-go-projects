package solver

import (
	"fmt"
	"image"
)

// Solver is capable of finding the path through a maze.
type Solver struct {
	maze *image.RGBA
}

// New builds a Solver by taking the path to the PNG maze.
func New(imagePath string) (*Solver, error) {
	img, err := openImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze: img,
	}, nil
}

// Solve finds the path from one end to the other.
func (s *Solver) Solve() error {
	return nil
}

// SaveSolution saves the image as a PNG file with the solution path in red.
func (s *Solver) SaveSolution(outputPath string) error {
	return nil
}
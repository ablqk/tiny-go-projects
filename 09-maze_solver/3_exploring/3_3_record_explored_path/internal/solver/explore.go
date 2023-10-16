package solver

import (
	"image"
	"log/slog"
)

// explore one path and publish to the s.pathsToExplore channel
// any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		// This is a safety net. It should be used, but when it's needed, at least it's there.
		return
	}

	pos := pathToBranch.at

	for {
		// We know we'll have up to 3 new neighbours to explore.
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				// Let's not return to the previous position
				continue
			}
			// Look at the colour of this pixel.
			// RGBAAt returns a color.RGBA{} zero value if the pixel is outside the bounds of the image.
			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				slog.Info("Treasure found!")
				return
			case s.config.pathColour:
				candidates = append(candidates, n)
			}
		}

		if len(candidates) == 0 {
			slog.Debug("I must have taken the wrong turn.", "position", pos)
			return
		}

		for _, candidate := range candidates[1:] {
			branch := &path{previousSteps: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		pathToBranch = &path{previousSteps: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}
}
package solver

import (
	"fmt"
	"image"
	"log/slog"
	"sync"
)

// listenToBranches creates a new routine for each branch published in s.pathsToExplore.
func (s *Solver) listenToBranches() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for {
		select {
		// s.quit will never return a value, unless something writes in it (which we don't do)
		// or it has been closed, which we do when we find the treasure.
		case <-s.quit:
			slog.Info("the treasure has been found, worker going to sleep")
			return
		case p := <-s.pathsToExplore:
			wg.Add(1)
			go func(p *path) {
				defer wg.Done()

				s.explore(p)
			}(p)
		}
	}
}

// explore one path and publish to the s.pathsToExplore channel
// any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		// This is a safety net. It should be used, but when it's needed, at least it's there.
		return
	}

	pos := pathToBranch.at

	for {
		// Let's first check whether we should quit.
		select {
		case <-s.quit:
			return
		default:
			// Continue the exploration.
		}

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
				s.mutex.Lock()
				if s.solution == nil {
					s.solution = &path{previousStep: pathToBranch, at: n}
					slog.Info(fmt.Sprintf("Treasure found: %v!", s.solution.at))
					close(s.quit)
				}
				s.mutex.Unlock()
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
			branch := &path{previousStep: pathToBranch, at: candidate}
			// We are sure we send to pathsToExplore only when the quit channel isn't closed.
			// A goroutine might have found the treasure since the check at the start of the loop.
			select {
			// s.quit returns a zero value only when the channel was closed, here -- when the exploration should end.
			case <-s.quit:
				slog.Debug("I'm an unlucky branch, someone else found the treasure, I give up.", "position", pos)
				return
			case s.pathsToExplore <- branch:
				// continue execution after the select block
			}
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}
}
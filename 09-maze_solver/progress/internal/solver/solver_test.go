package solver

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolver_findEntrance_success(t *testing.T) {
	tests := map[string]struct {
		inputPath string
		want      image.Point
	}{
		"middle": {
			inputPath: "testdata/maze10_10.png",
			want:      image.Point{X: 0, Y: 5},
		},
		"400 px": {
			inputPath: "testdata/maze400_400.png",
			want:      image.Point{X: 0, Y: 200},
		},
		"treasure near entrance": {
			inputPath: "testdata/maze10_exit.png",
			want:      image.Point{X: 0, Y: 5},
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			img, err := openMaze(tt.inputPath)
			require.NoError(t, err)

			s := &Solver{
				maze:   img,
				config: defaultColours(),
			}

			got, err := s.findEntrance()

			assert.Equalf(t, tt.want, got, "findEntrance()")
		})
	}
}

func TestSolver_findEntrance_error(t *testing.T) {
	tests := map[string]struct {
		inputPath string
	}{
		"entrance in a corner": {
			inputPath: "testdata/maze10_corner.png",
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			img, err := openMaze(tt.inputPath)
			require.NoError(t, err)

			s := &Solver{
				maze:   img,
				config: defaultColours(),
			}

			_, err = s.findEntrance()

			assert.Error(t, err)
		})
	}
}

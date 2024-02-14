//go:build exercise

package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	const layout = "02 January 2006"

	tests := map[string]struct {
		include time.Time
		want    FormattedWeek
	}{
		"Mon. 01 Jan 2024": {
			include: time.Date(2024, 1, 1, 12, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 1, 7, 23, 59, 0, 0, time.UTC),
			},
		},
		"Wed. 14 Feb. 2024": {
			include: time.Date(2024, 2, 14, 15, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2024, 2, 12, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2024, 2, 18, 23, 59, 0, 0, time.UTC),
			},
		},
		"Sun. 01 Jan 2023": {
			include: time.Date(2023, 1, 1, 5, 54, 23, 2, time.UTC),
			want: FormattedWeek{
				start: time.Date(2022, 12, 26, 0, 0, 0, 0, time.UTC),
				end:   time.Date(2023, 1, 1, 23, 59, 0, 0, time.UTC),
			},
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := NewWeek(tt.include, layout)
			assert.Equal(t, tt.want.start.String(), got.start.String())
			assert.Equal(t, tt.want.end.String(), got.end.String())
		})
	}
}
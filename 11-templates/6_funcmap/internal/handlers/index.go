package handlers

import (
	_ "embed"
	"html/template"
	"net/http"
	"strings"
	"time"

	"learngo-pockets/templates/internal/habit"
)

//go:embed index.html
var indexPage string

// index serves the root page of the app.
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	const indexEndpoint = "index"

	habits, err := s.client.ListHabits(r.Context(), time.Now())
	if err != nil {
		s.logAndHideError(w, indexEndpoint, err, http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("index").
		Funcs(template.FuncMap{"scoreStatus": scoreStatus, "progress": progress}).
		Parse(indexPage)
	if err != nil {
		s.logAndHideError(w, indexEndpoint, err, http.StatusInternalServerError)
		return
	}

	week := habit.NewWeek(time.Now(), "02 January 2006")

	err = tpl.Execute(w, map[string]interface{}{
		"Habits": habits,
		"Date":   week,
	})
	if err != nil {
		s.logAndHideError(w, indexEndpoint, err, http.StatusInternalServerError)
		return
	}
}

func scoreStatus(habit *habit.Habit) string {
	prog := float32(habit.Ticks) / float32(habit.WeeklyFrequency)
	switch {
	case prog == 0:
		return "not_started"
	case prog < 0.5:
		return "started"
	case prog < 1:
		return "good_progress"
	default:
		return "completed"
	}
}

func progress(habit *habit.Habit) string {
	prog := min(int(float32(habit.Ticks)/float32(habit.WeeklyFrequency)*10), 10)
	return strings.Repeat("#", prog) + strings.Repeat("_", 10-prog)
}

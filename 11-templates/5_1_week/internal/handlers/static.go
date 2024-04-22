package handlers

import (
	"net/http"
	"slices"

	"github.com/go-chi/chi/v5"
)

// index serves the root page of the app.
func (s *Server) assets(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "filename")

	// prevent injection
	if !isValidName(fileName) {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeFile(w, r, "internal/assets/"+fileName)
}

func isValidName(name string) bool {
	return slices.Contains([]string{
		"styles.css",
	}, name)
}

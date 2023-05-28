package handlers

import (
	"database/sql"
	"errors"
	"net/http"
)

func (h *Handler) CheckLink(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Query().Get("link")
	code, err := h.service.Check(link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Redirect(w, r, r.URL.Path, http.StatusMovedPermanently)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write([]byte("Success"))
}

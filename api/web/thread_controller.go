package web

import (
	"api"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)


func (h *Handler) ThreadList() http.HandlerFunc {
	type data struct {
		Thread []api.Thread
	} 

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(tt)

	}
}


func (h *Handler) ThreadById() http.HandlerFunc {
	type data struct {
		Thread []api.Thread
	} 

	//tmpl := template.Must(template.New("").Parse(threadsListHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// id, _ := uuid.Parse(r.URL.Query().Get("id"))

 		id, _ := uuid.Parse(chi.URLParam(r, "id"))

		tt, err := h.store.Thread(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(tt)

	}
}
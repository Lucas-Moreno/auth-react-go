package web

import (
	"api"
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"fmt"
)


func (h *Handler) ThreadList() http.HandlerFunc {
	type data struct {
		Thread []api.Thread
	} 

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tt, err := h.store.Threads()
		fmt.Println(tt)
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

func (h *Handler) CreateThread() http.HandlerFunc {
	type data struct {
		Thread api.Thread
	} 

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		i := api.Thread{}

		_ = json.NewDecoder(r.Body).Decode(&i)

		if err := h.store.CreateThread(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_  = json.NewEncoder(w).Encode(struct{
			Message string `json:"message"`
		}{
			Message: "Thread created",
		})
	}
}

func (h *Handler) UpdateThread() http.HandlerFunc {
	type data struct {
		Thread api.Thread
	} 

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		i := api.Thread{}
		id, _ := uuid.Parse(chi.URLParam(r, "id"))

		thread, _ := h.store.Thread(id)

		_ = json.NewDecoder(r.Body).Decode(&i)

		if i.Title != "" {
			thread.Title = i.Title
		}
		if i.Description != "" {
			thread.Description = i.Description
		}


		if err := h.store.UpdateThread(thread); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_  = json.NewEncoder(w).Encode(struct{
			Message string `json:"message"`
		}{
			Message: "Thread updated",
		})
	}
}
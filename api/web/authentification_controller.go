package web


import (
	"api"
	"net/http"
	"encoding/json"
)


func (h *Handler) SignUp() http.HandlerFunc {
	type data struct {
		User api.User
	} 

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		i := api.User{}

		_ = json.NewDecoder(r.Body).Decode(&i)

		if err := h.store.SignUp(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_  = json.NewEncoder(w).Encode(struct{
			Message string `json:"message"`
		}{
			Message: "User created",
		})
	}
}

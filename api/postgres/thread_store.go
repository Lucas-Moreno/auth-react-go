package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"fmt"
	"api"
)

type ThreadStore struct {
	*sqlx.DB
}

func (s *ThreadStore) Thread(id uuid.UUID) (api.Thread, error) {
	var t api.Thread
	if err := s.Get(&t, `SELECT * FROM threads WHERE id=$1`, id); err != nil {
		return api.Thread{}, fmt.Errorf("Error getting Thread: %w", err)
	}
	return t, nil
}
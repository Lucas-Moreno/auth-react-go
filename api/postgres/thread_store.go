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

func (s *ThreadStore) Threads() ([]api.Thread, error) {
	var t []api.Thread
	if err := s.Select(&t, `SELECT * FROM threads`); err != nil {
		return []api.Thread{}, fmt.Errorf("Error getting Threads: %w", err)
	}
	return t, nil
}

func (s *ThreadStore) CreateThread(t api.Thread) error {
	if err := s.Get(&t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`, 
		t.ID,
		t.Title,
		t.Description); err != nil {
		return fmt.Errorf("Error create Thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t api.Thread) error {
	if err := s.Get(&t, `UPDATE threads SET title = $1, description = $2 WHERE id=$3 RETURNING *`, 
		t.Title,
		t.Description,
		t.ID); err != nil {
		return fmt.Errorf("Error create Thread: %w", err)
	}
	return nil
}
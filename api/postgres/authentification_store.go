package postgres

import (
	"github.com/jmoiron/sqlx"
	"api"
	"fmt"
	_ "github.com/lib/pq"
)

type AuthentificationStore struct {
	*sqlx.DB
}


func (s *AuthentificationStore) SignUp(t api.User) error {
	if err := s.Get(&t, `INSERT INTO users (ID, FirstName, LastName, Password) VALUES ($1, $2, $3, $4) RETURNING *`, 
		t.ID,
		t.FirstName,
		t.LastName,
		t.Password); err != nil {
		return fmt.Errorf("Error create User: %w", err)
	}
	return nil
}
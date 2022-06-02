package api

import (
	"github.com/google/uuid"
)

type Thread struct {
	ID				uuid.UUID	`db:"id"`
	Title			string		`db:"title"`
	Description		string		`db:"description"`
	UserId			uuid.UUID	`db:"user_id"`

}

type User struct {
	ID				uuid.UUID	`db:"id"`
	FirstName		string		`db:"firstname"`
	LastName		string		`db:"lastname"`
	Password		string		`db:"password"`
}

type ThreadStore interface {
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error)
	CreateThread(t Thread) error
	UpdateThread(t Thread) error
	DeleteThread(id uuid.UUID) error
}


type AuthentificationStore interface {
	SignUp(t User) error
}

type Store interface {
	ThreadStore
	AuthentificationStore
}



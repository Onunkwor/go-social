package store

import (
	"database/sql"
)

type Storage struct {
	Post  *PostStore
	Users *UsersStore
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:  &PostStore{db: db},
		Users: &UsersStore{db: db},
	}
}

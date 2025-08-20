package store

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Storage struct {
	Post  *PostStore
	Users *UsersStore
}

func NewStorage(db *mongo.Database) Storage {
	return Storage{
		Post:  &PostStore{db: db},
		Users: &UsersStore{db: db},
	}
}

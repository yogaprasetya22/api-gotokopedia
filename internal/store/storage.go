package store

import "database/sql"

type Storage struct {
	Products interface{}
	// Users interface{}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Products: &ProductStore{db},
		// Users: NewUserStore(db),
	}
}

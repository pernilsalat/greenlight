package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type CRUD[T any] interface {
	Insert(movie *T) error
	Get(id int64) (*T, error)
	Update(movie *T) error
	Delete(id int64) error
}

type Models struct {
	Movies CRUD[Movie]
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: &MovieModel{DB: db},
	}
}

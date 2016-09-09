package storage

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

// DB provides basic database operations
// (implemented by scribble db)
type DB interface {
	Write(collection, resource string, v interface{}) error
	Read(collection, resource string, v interface{}) error
	ReadAll(collection string) ([]string, error)
	Delete(collection, resource string) error
}

func NewDB(dir string) (DB, error) {
	db, err := scribble.New(dir, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating db: %v", err)
	}

	return db, nil
}

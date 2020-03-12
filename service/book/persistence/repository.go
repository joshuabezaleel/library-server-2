package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importing postgreSQL driver

	"github.com/joshuabezaleel/library-server-2/service/book/book"
)

// Repository holds dependencies for the current persistence layer.
type Repository struct {
	BookRepository book.Repository

	DB *sql.DB
}

// NewRepository returns a new Repository
// with all of the necessary dependencies.
func NewRepository() *Repository {

	connectionParameters := fmt.Sprintf("host=localhost port=8081 user=postgres password=postgres dbname=library-server-2 sslmode=disable")
	DB, err := sql.Open("postgres", connectionParameters)
	if err != nil {
		panic(err)
	}

	bookRepository := NewBookRepository(DB)

	repository := &Repository{
		BookRepository: bookRepository,
		DB:             DB,
	}

	return repository
}

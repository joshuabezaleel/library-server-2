package persistence

import (
	"database/sql"

	"github.com/joshuabezaleel/library-server-2/service/book/book"
)

type bookRepository struct {
	DB *sql.DB
}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository(DB *sql.DB) book.Repository {
	return &bookRepository{
		DB: DB,
	}
}

func (repo *bookRepository) GetAll() ([]*book.Book, error) {
	var oneBook *book.Book
	var allBooks []*book.Book

	rows, err := repo.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(oneBook)
		if err != nil {
			return nil, err
		}

		allBooks = append(allBooks, oneBook)
	}

	return allBooks, nil
}

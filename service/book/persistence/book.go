package persistence

import (
	"github.com/joshuabezaleel/library-server-2/service/book/book"
)

var allBooks = []*book.Book{
	&book.Book{ID: 1, Title: "", Description: ""},
	&book.Book{ID: 2, Title: "", Description: ""},
	&book.Book{ID: 3, Title: "", Description: ""},
	&book.Book{ID: 4, Title: "", Description: ""},
	&book.Book{ID: 5, Title: "", Description: ""},
	&book.Book{ID: 6, Title: "", Description: ""},
}

type bookRepository struct{}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository() book.Repository {
	return &bookRepository{}
}

func (repo *bookRepository) GetAll() ([]*book.Book, error) {
	// var oneBook *book.Book
	// var allBooks []*book.Book

	// rows, err := repo.DB.Query("SELECT * FROM books")
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	err = rows.Scan(oneBook)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	allBooks = append(allBooks, oneBook)
	// }

	// return allBooks, nil
	return allBooks, nil
}

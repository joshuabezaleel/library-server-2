package main

import (
	"github.com/joshuabezaleel/library-server-2/service/book/book"
	"github.com/joshuabezaleel/library-server-2/service/book/persistence"
	"github.com/joshuabezaleel/library-server-2/service/book/server"
)

func main() {
	repository := persistence.NewRepository()
	defer repository.DB.Close()

	bookService := book.NewBookService(repository.BookRepository)

	srv := server.NewServer(bookService)
	srv.Run("8082")
}

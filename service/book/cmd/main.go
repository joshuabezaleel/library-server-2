package main

import (
	"github.com/joshuabezaleel/library-server-2/service/book/book"
	"github.com/joshuabezaleel/library-server-2/service/book/persistence"
	"github.com/joshuabezaleel/library-server-2/service/book/server"
)

func main() {
	bookRepository := persistence.NewBookRepository()

	bookService := book.NewBookService(bookRepository)

	srv := server.NewServer(bookService)
	srv.Run("8082")
}

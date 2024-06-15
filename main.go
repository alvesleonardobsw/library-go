package main

import (
	"github.com/alvesleonardobsw/library-go/controller"
	"github.com/alvesleonardobsw/library-go/domain/author"
	"github.com/alvesleonardobsw/library-go/domain/book"
)

func main() {
	authorService := author.NewAuthorService()
	bookService := book.NewBookService()

	r := controller.NewRouter(authorService, bookService)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

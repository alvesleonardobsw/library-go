package main

import (
	"log"

	"github.com/alvesleonardobsw/library-go/controller"
	"github.com/alvesleonardobsw/library-go/domain/author"
	"github.com/alvesleonardobsw/library-go/domain/book"
	"github.com/alvesleonardobsw/library-go/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=library sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	authorRepository := repository.NewAuthorRepository(db)
	bookRepository := repository.NewBookRepository(db)

	authorService := author.NewAuthorService(authorRepository)
	bookService := book.NewBookService(bookRepository)

	r := controller.NewRouter(authorService, bookService)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

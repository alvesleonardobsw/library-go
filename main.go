package main

import (
	"github.com/alvesleonardobsw/library-go/controller"
	"github.com/alvesleonardobsw/library-go/domain/author"
)

func main() {
	authorService := author.NewAuthorService()
	r := controller.NewRouter(authorService)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

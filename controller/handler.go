package controller

import (
	"github.com/alvesleonardobsw/library-go/domain/author"
	"github.com/alvesleonardobsw/library-go/domain/book"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authorService *author.AuthorService
	bookService   *book.BookService
}

func NewRouter(authorService *author.AuthorService, bookService *book.BookService) *gin.Engine {
	r := gin.Default()

	h := Handler{
		authorService: authorService,
		bookService:   bookService,
	}

	//author
	r.GET("/authors", h.GetAllAuthors)
	r.POST("/author", h.PostAuthor)
	r.PUT("/author/:authorId", h.PutAuthor)
	r.DELETE("/author/:authorId", h.DeleteAuthor)

	//book
	r.GET("/books", h.GetAllBooks)
	r.POST("/book", h.PostBook)
	r.PUT("/book/:bookId", h.PutBook)
	r.DELETE("/book/:bookId", h.DeleteBook)

	return r
}

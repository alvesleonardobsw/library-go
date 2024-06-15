package controller

import (
	"github.com/alvesleonardobsw/library-go/domain/author"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authorService *author.AuthorService
}

func NewRouter(authorService *author.AuthorService) *gin.Engine {
	r := gin.Default()

	h := Handler{
		authorService: authorService,
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

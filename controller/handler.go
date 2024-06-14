package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//author
	r.GET("/authors", GetAllAuthors)
	r.POST("/author", PostAuthor)
	r.PUT("/author/:authorId", PutAuthor)
	r.DELETE("/author/:authorId", DeleteAuthor)

	//book
	r.GET("/books", GetAllBooks)
	r.POST("/book", PostBook)
	r.PUT("/book/:bookId", PutBook)
	r.DELETE("/book/:bookId", DeleteBook)

	return r
}

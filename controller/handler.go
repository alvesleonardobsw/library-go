package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//author
	r.GET("/authors", GetAllAuthors)
	r.POST("/author", PostAuthor)

	//book

	return r
}

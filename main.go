package main

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
)

var authorsList = make([]domain.Author, 0)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/authors", func(c *gin.Context) {
		c.JSON(http.StatusOK, authorsList)
	})

	r.POST("/author", func(c *gin.Context) {
		var newAuthor domain.Author

		if err := c.BindJSON(&newAuthor); err != nil {
			fmt.Println(err.Error())
			c.Status(http.StatusInternalServerError)
			return
		}

		authorsList = append(authorsList, newAuthor)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

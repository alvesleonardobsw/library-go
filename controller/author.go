package controller

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
)

var authorsList = make([]domain.Author, 0)

func GetAllAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authorsList)
}

func PostAuthor(c *gin.Context) {
	var newAuthor domain.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	authorsList = append(authorsList, newAuthor)
}

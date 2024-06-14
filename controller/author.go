package controller

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	newAuthor.Id = uuid.New().String()
	authorsList = append(authorsList, newAuthor)
}

func PutAuthor(c *gin.Context) {
	authorId := c.Param("authorId")
	var editedAuthor domain.Author

	if err := c.BindJSON(&editedAuthor); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(authorsList); i++ {
		if authorId == authorsList[i].Id {
			authorsList[i].Name = editedAuthor.Name
			authorsList[i].LastName = editedAuthor.LastName
			c.JSON(http.StatusOK, authorsList[i])
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func DeleteAuthor(c *gin.Context) {
	authorId := c.Param("authorId")

	for i := 0; i < len(authorsList); i++ {
		if authorId == authorsList[i].Id {
			authorsList = append(authorsList[:i], authorsList[i+1:]...)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

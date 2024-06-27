package controller

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllAuthors(c *gin.Context) {

	authorsSearch := []string{}
	name := c.Query("name")
	lastname := c.Query("lastName")
	authorsSearch = append(authorsSearch, name, lastname)

	authors, err := h.authorService.GetAll(authorsSearch)
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, authors)
}

func (h *Handler) PostAuthor(c *gin.Context) {
	var newAuthor domain.Author

	if err := c.BindJSON(&newAuthor); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := h.authorService.Create(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": newAuthor.Id})
}

func (h *Handler) PutAuthor(c *gin.Context) {
	authorId := c.Param("authorId")
	var editedAuthor domain.Author

	if err := c.BindJSON(&editedAuthor); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	author, err := h.authorService.Update(authorId, &editedAuthor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) DeleteAuthor(c *gin.Context) {
	authorId := c.Param("authorId")

	if err := h.authorService.Delete(authorId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

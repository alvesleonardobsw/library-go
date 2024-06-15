package controller

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var booksList = make([]domain.Book, 0)

func (h *Handler) GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, booksList)
}

func (h *Handler) PostBook(c *gin.Context) {
	var newBook domain.Book

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	newBook.Id = uuid.New().String()
	booksList = append(booksList, newBook)
}

func (h *Handler) PutBook(c *gin.Context) {
	bookId := c.Param("bookId")
	var editedBook domain.Book

	if err := c.BindJSON(&editedBook); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(booksList); i++ {
		if bookId == booksList[i].Id {
			booksList[i].Name = editedBook.Name
			booksList[i].Genre = editedBook.Genre
			c.JSON(http.StatusOK, booksList[i])
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")

	for i := 0; i < len(booksList); i++ {
		if bookId == booksList[i].Id {
			booksList = append(booksList[:i], booksList[i+1:]...)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

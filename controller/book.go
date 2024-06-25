package controller

import (
	"fmt"
	"net/http"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAll()
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *Handler) PostBook(c *gin.Context) {
	var newBook domain.Book

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := h.bookService.Create(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": newBook.Id})
}

func (h *Handler) PutBook(c *gin.Context) {
	bookId := c.Param("bookId")
	var editedBook domain.Book

	if err := c.BindJSON(&editedBook); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	book, err := h.bookService.Update(bookId, &editedBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")

	if err := h.bookService.Delete(bookId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

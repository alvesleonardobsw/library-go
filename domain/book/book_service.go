package book

import (
	"fmt"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/google/uuid"
)

type BookService struct {
}

var booksList = make([]domain.Book, 0)

func NewBookService() *BookService {
	return &BookService{}
}

func (as *BookService) Create(book *domain.Book) error {
	book.Id = uuid.New().String()

	if err := book.ValidateBook(); err != nil {
		return err
	}

	booksList = append(booksList, *book)

	return nil
}

func (as *BookService) GetAll() []domain.Book {
	return booksList
}

func (as *BookService) Update(bookId string, editedBook domain.Book) (*domain.Book, error) {

	if err := editedBook.ValidateBook(); err != nil {
		return nil, err
	}

	for i := 0; i < len(booksList); i++ {
		if bookId == booksList[i].Id {
			booksList[i].Name = editedBook.Name
			booksList[i].Genre = editedBook.Genre
			return &booksList[i], nil
		}
	}
	return nil, fmt.Errorf("could not found Book with id %s", bookId)
}

func (as *BookService) Delete(bookId string) error {
	for i := 0; i < len(booksList); i++ {
		if bookId == booksList[i].Id {
			booksList = append(booksList[:i], booksList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not found book with id %s", bookId)
}

package book

import (
	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/alvesleonardobsw/library-go/repository"
	"github.com/google/uuid"
)

type BookService struct {
	repository *repository.BookRepository
}

func NewBookService(repository *repository.BookRepository) *BookService {
	return &BookService{
		repository: repository,
	}
}

func (as *BookService) Create(book *domain.Book) error {
	book.Id = uuid.New().String()

	if err := book.ValidateBook(); err != nil {
		return err
	}

	return as.repository.Insert(book)
}

func (as *BookService) GetAll() ([]domain.Book, error) {
	return as.repository.FindAll()
}

func (as *BookService) Update(bookId string, editedBook *domain.Book) (*domain.Book, error) {

	if err := editedBook.ValidateBook(); err != nil {
		return nil, err
	}

	if err := as.repository.Update(bookId, editedBook); err != nil {
		return nil, err
	}

	return editedBook, nil
}

func (as *BookService) Delete(bookId string) error {
	if err := as.repository.Delete(bookId); err != nil {
		return err
	}

	return nil
}

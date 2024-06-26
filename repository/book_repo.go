package repository

import (
	"fmt"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (ar *BookRepository) FindAll() ([]domain.Book, error) {
	Books := []domain.Book{}

	if err := ar.db.Select(&Books, "SELECT * FROM Book"); err != nil {
		return nil, err
	}

	return Books, nil
}

func (ar *BookRepository) Insert(book *domain.Book) error {
	_, err := ar.db.NamedExec("INSERT INTO Book (IdBook, Name, Genre, IdAuthor) VALUES (:idbook, :name, :genre, :idauthor)", book)

	return err
}

func (ar *BookRepository) Update(bookId string, editedBook *domain.Book) error {
	editedBook.Id = bookId
	result, err := ar.db.NamedExec("UPDATE Book SET Name=:name, Genre=:genre WHERE IdBook = :idbook", editedBook)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("could not found book with id %s", bookId)
	}

	return nil
}

func (ar *BookRepository) Delete(bookId string) error {
	result, err := ar.db.Exec("DELETE FROM Book WHERE IdBook = $1", bookId)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("could not found book with id %s", bookId)
	}

	return nil
}

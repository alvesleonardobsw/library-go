package repository

import (
	"fmt"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/jmoiron/sqlx"
)

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (ar *AuthorRepository) FindAll(authorsSearch []string) ([]domain.Author, error) {
	authors := []domain.Author{}

	query := "SELECT * FROM Author"

	for i := 0; i < len(authorsSearch)-1; i++ {
		if authorsSearch[i] != "" || authorsSearch[i+1] != "" {
			query += fmt.Sprintf(" WHERE name ILIKE '%%%s%%' and lastName ILIKE'%%%s%%'", authorsSearch[i], authorsSearch[i+1])
			break
		}
	}

	if err := ar.db.Select(&authors, query); err != nil {
		return nil, err
	}

	return authors, nil
}

func (ar *AuthorRepository) Insert(author *domain.Author) error {
	_, err := ar.db.NamedExec("INSERT INTO Author (IdAuthor, Name, LastName) VALUES (:idauthor, :name, :lastname)", author)

	return err
}

func (ar *AuthorRepository) Update(authorId string, editedAuthor *domain.Author) error {
	editedAuthor.Id = authorId
	result, err := ar.db.NamedExec("UPDATE Author SET Name=:name, LastName=:lastname WHERE IdAuthor = :idauthor", editedAuthor)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("could not found author with id %s", authorId)
	}

	return nil
}

func (ar *AuthorRepository) Delete(authorId string) error {
	result, err := ar.db.Exec("DELETE FROM Author WHERE IdAuthor = $1", authorId)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("could not found author with id %s", authorId)
	}

	return nil
}

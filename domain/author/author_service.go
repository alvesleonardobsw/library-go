package author

import (
	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/alvesleonardobsw/library-go/repository"
	"github.com/google/uuid"
)

type AuthorService struct {
	repository *repository.AuthorRepository
}

func NewAuthorService(repository *repository.AuthorRepository) *AuthorService {
	return &AuthorService{
		repository: repository,
	}
}

func (as *AuthorService) Create(author *domain.Author) error {
	author.Id = uuid.New().String()

	if err := author.Validate(); err != nil {
		return err
	}

	return as.repository.Insert(author)
}

func (as *AuthorService) GetAll(authorsSearch []string) ([]domain.Author, error) {
	return as.repository.FindAll(authorsSearch)
}

func (as *AuthorService) Update(authorId string, editedAuthor *domain.Author) (*domain.Author, error) {
	if err := editedAuthor.Validate(); err != nil {
		return nil, err
	}

	if err := as.repository.Update(authorId, editedAuthor); err != nil {
		return nil, err
	}

	return editedAuthor, nil
}

func (as *AuthorService) Delete(authorId string) error {
	if err := as.repository.Delete(authorId); err != nil {
		return err
	}

	return nil
}

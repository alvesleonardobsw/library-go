package author

import (
	"fmt"

	"github.com/alvesleonardobsw/library-go/domain"
	"github.com/google/uuid"
)

type AuthorService struct {
}

var authorsList = make([]domain.Author, 0)

func NewAuthorService() *AuthorService {
	return &AuthorService{}
}

func (as *AuthorService) Create(author *domain.Author) error {
	author.Id = uuid.New().String()

	if err := author.Validate(); err != nil {
		return err
	}

	authorsList = append(authorsList, *author)

	return nil
}

func (as *AuthorService) GetAll() []domain.Author {
	return authorsList
}

func (as *AuthorService) Update(authorId string, editedAuthor domain.Author) (*domain.Author, error) {

	if err := editedAuthor.Validate(); err != nil {
		return nil, err
	}

	for i := 0; i < len(authorsList); i++ {
		if authorId == authorsList[i].Id {
			authorsList[i].Name = editedAuthor.Name
			authorsList[i].LastName = editedAuthor.LastName
			return &authorsList[i], nil
		}
	}
	return nil, fmt.Errorf("could not found author with id %s", authorId)
}

func (as *AuthorService) Delete(authorId string) error {
	for i := 0; i < len(authorsList); i++ {
		if authorId == authorsList[i].Id {
			authorsList = append(authorsList[:i], authorsList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not found author with id %s", authorId)
}

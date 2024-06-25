package domain

import "fmt"

type Book struct {
	Name  string `json:"name" db:"name"`
	Genre string `json:"genre" db:"genre"`
	Id    string `json:"id" db:"idbook"`
}

func (a *Book) ValidateBook() error {
	if a.Name == "" {
		return fmt.Errorf("name of book cannot be blank")
	}

	if a.Genre == "" {
		return fmt.Errorf("genre of book cannot be blank")
	}
	return nil
}

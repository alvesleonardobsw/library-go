package domain

import "fmt"

type Book struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Id    string `json:"id"`
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

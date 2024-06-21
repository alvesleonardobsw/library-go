package domain

import "fmt"

type Author struct {
	Name     string `json:"name" db:"name"`
	LastName string `json:"lastName" db:"lastname"`
	Id       string `json:"id" db:"idauthor"`
}

func (a *Author) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("name of author cannot be blank")
	}

	if a.LastName == "" {
		return fmt.Errorf("lastname of author cannot be blank")
	}
	return nil
}

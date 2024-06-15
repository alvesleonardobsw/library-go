package domain

import "fmt"

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Id       string `json:"id"`
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

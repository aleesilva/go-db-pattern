package entity

import uuid "github.com/satori/go.uuid"

type PersonRepository interface {
	Get(id string) (Person, error)
	Create(person Person) (Person, error)
}

type Person struct {
	ID    string
	Name  string
	Email string
}

func NewPerson() *Person {
	person := Person{
		ID: uuid.NewV4().String(),
	}
	return &person
}

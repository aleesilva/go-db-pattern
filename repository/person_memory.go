package repository

import (
	"errors"

	"github.com/aleesilva/go-db-pattern/entity"
)

type PersonsMemoryDB struct {
	Persons []entity.Person
}

type PersonRepositoryMemory struct {
	db PersonsMemoryDB
}

func NewPersonRepositoryMemory(db PersonsMemoryDB) *PersonRepositoryMemory {
	return &PersonRepositoryMemory{db: db}
}

func (p *PersonRepositoryMemory) Get(id string) (entity.Person, error) {
	for _, person := range p.db.Persons {
		if person.ID == id {
			return person, nil
		}
	}
	return entity.Person{}, errors.New("not found")
}

func (p *PersonRepositoryMemory) Create(person entity.Person) (entity.Person, error) {
	p.db.Persons = append(p.db.Persons, person)
	return person, nil
}

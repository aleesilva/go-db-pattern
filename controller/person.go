package controller

import "github.com/aleesilva/go-db-pattern/entity"

type PersonController struct {
	Repository entity.PersonRepository
}

func NewPersonController(repository entity.PersonRepository) *PersonController {
	return &PersonController{Repository: repository}
}

func (p *PersonController) FindById(id string) (entity.Person, error) {
	person, error := p.Repository.Get(id)

	if error != nil {
		return entity.Person{}, error
	}

	return person, nil
}

func (p *PersonController) Create(name string, email string) (entity.Person, error) {
	person := entity.NewPerson()
	person.Name = name
	person.Email = email

	newPerson, err := p.Repository.Create(*person)

	if err != nil {
		return entity.Person{}, err
	}

	return newPerson, nil
}

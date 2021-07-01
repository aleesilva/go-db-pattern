package repository

import (
	"database/sql"

	"github.com/aleesilva/go-db-pattern/entity"

	_ "github.com/lib/pq"
)

type PersonRepositoryPostgres struct {
	db *sql.DB
}

func NewPersonRepositoryPostgres(db *sql.DB) *PersonRepositoryPostgres {
	return &PersonRepositoryPostgres{db: db}
}

func (p *PersonRepositoryPostgres) Get(id string) (entity.Person, error) {
	var person entity.Person

	stmt, err := p.db.Prepare("select id, name, email from persons where id=?")

	if err != nil {
		return entity.Person{}, err
	}

	err = stmt.QueryRow(id).Scan(&person.ID, &person.Name, &person.Email)

	if err != nil {
		return entity.Person{}, err
	}

	return person, nil
}

func (p *PersonRepositoryPostgres) Create(person entity.Person) (entity.Person, error) {
	stmt, err := p.db.Prepare(`insert into persons (id, name, email) values ($1 , $2 , $3)`)

	if err != nil {
		return entity.Person{}, err
	}

	_, err = stmt.Exec(person.ID, person.Name, person.Email)

	if err != nil {
		return entity.Person{}, err
	}

	err = stmt.Close()

	if err != nil {
		return entity.Person{}, err
	}

	return person, nil
}

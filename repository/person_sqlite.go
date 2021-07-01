package repository

import (
	"database/sql"

	"github.com/aleesilva/go-db-pattern/entity"

	_ "github.com/mattn/go-sqlite3"
)

type PersonRepositorySqlite struct {
	db *sql.DB
}

func NewPersonSqlite(db *sql.DB) *PersonRepositorySqlite {
	return &PersonRepositorySqlite{db: db}
}

func (p *PersonRepositorySqlite) Get(id string) (entity.Person, error) {
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

func (p *PersonRepositorySqlite) Create(person entity.Person) (entity.Person, error) {
	stmt, err := p.db.Prepare(`insert into persons (id, name, email) values(?, ? , ? ) `)

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

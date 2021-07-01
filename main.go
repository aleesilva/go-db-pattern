package main

import (
	"database/sql"
	"fmt"

	"github.com/aleesilva/go-db-pattern/controller"
	"github.com/aleesilva/go-db-pattern/repository"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "golang-db"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disabled", host, port, user, password, dbname)
	// fmt.Println(psqlconn)
	// db := repository.PersonsMemoryDB{Persons: []entity.Person{}}
	// repositoryMemory := repository.NewPersonRepositoryMemory(db)
	// db, _ := sql.Open("sqlite3", "./sqlite.repository")
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// repositorySqlite := repository.NewPersonSqlite(db)
	repositoryPostgres := repository.NewPersonRepositoryPostgres(db)
	ct := controller.NewPersonController(repositoryPostgres)

	persons, _ := ct.Create("teste user", "teste@teste.com.br")

	fmt.Println(persons)
}

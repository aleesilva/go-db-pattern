package main

import (
	"database/sql"
	"fmt"

	"github.com/aleesilva/go-db-pattern/controller"
	"github.com/aleesilva/go-db-pattern/repository"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "golangdb"
)

func Error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// fmt.Println(psqlconn)
	// db := repository.PersonsMemoryDB{Persons: []entity.Person{}}
	// repositoryMemory := repository.NewPersonRepositoryMemory(db)
	// db, _ := sql.Open("sqlite3", "./sqlite.repository")
	db, err := openDB("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// repositorySqlite := repository.NewPersonSqlite(db)
	repositoryPostgres := repository.NewPersonRepositoryPostgres(db)
	ct := controller.NewPersonController(repositoryPostgres)

	person, err := ct.Create("teste user", "teste@teste.com.br")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %T", person, person)
}

func openDB(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

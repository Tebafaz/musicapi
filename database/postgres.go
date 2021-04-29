package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	//postgres driver
	_ "github.com/lib/pq"
)

//Postgres a
var postgres *sqlx.DB

//InitPostgres a
func InitPostgres() {
	var err error
	postgres, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode))
	if err != nil {
		panic(err)
	}
	fmt.Printf("connection to %s created\n", postgres.DriverName())
}

//ClosePostgres a
func ClosePostgres() {
	err := postgres.Close()
	if err != nil {
		panic("postgres refused to close")
	}
}

package users_db

import (
	"database/sql"
	"fmt"
	"log"

	// _ means import but we don't need it in the code
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		"root",
		"goroot",
		"127.0.0.1",
		"3307",
		"users_db",
	)
	var err error
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Print("Database successfully configured.")
}

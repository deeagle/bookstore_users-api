package users_db

import (
	"database/sql"
	"fmt"
	"log"

	// _ means import but we don't need it in the code
	_ "github.com/go-sql-driver/mysql"
)

const (
// trainer uses not handy stuff (for me)
// environment variables (export) instead of a config.yml (chapter 3:10)
)

var (
	Client *sql.DB

	// trainer uses not handy stuff (for me)
	// environment variables (export) instead of a config.yml (chapter 3:10)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		"root",
		"goroot",
		"127.0.0.1",
		"3307",
		"users_db",
	)
	fmt.Println(datasourceName)
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

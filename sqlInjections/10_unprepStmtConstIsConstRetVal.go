package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtConstIsConstRetVal() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/gophersAreLovely")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := source()

	// Fetching Data from a Database
	rows, err := db.Query("SELECT * FROM users WHERE id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func source() string {
	return "I am an evil gopher"
}

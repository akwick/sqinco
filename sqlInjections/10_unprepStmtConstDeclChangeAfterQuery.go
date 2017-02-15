package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtConstDeclChangeAfterQuery() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := "1"

	// Fetching Data from a Database
	// @ExpectWarning false
	rows, err := db.Query("SELECT * FROM users WHERE id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Change s so that the value is not constant for a whole program analysis
	s = "2"
}

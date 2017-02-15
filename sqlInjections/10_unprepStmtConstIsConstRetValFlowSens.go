package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtConstIsConstRetValFlowSens() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/gophersAreLovely")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := source()
	s = "Lovely Gopher"

	// Fetching Data from a Database
	// @ExpectWarning false
	rows, err := db.Query("SELECT * FROM users WHERE id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

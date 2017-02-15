package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func WrappedFunc() {
	// Accessing a DB
	db, err = sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/gophersAreLovely")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const s = "Lovely Gopher"

	// Fetching Data from a Database
	// @ExpectWarning false
	rows, err := WrappedQuery("SELECT * FROM users WHERE id =", s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func WrappedQuery(s string, args interface{}) (*sql.Rows, error) {
	return db.Query(s, args)
}

package sqlInjections

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtArg() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fetching Data from a Database
	rows, err := db.Query("SELECT * FROM users WHERE id =" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

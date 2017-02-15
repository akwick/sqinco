package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtDeclStringQueryRow() {
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
	_ = db.QueryRow("SELECT * FROM users WHERE id =" + s)
}

package sqlInjections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtConstDeclSmallKeywords() {
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
	rows, err := db.Query("select * from users where id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

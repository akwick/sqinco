package sqlInjections

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtFormatedString() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/gophersAreLovely")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const1 := "SELECT * FROM foo WHERE name = '%s'"
	const2 := "Lovely Gopher"
	s := fmt.Sprintf(const1, const2)

	// Fetching Data from a Database
	// @ExpectWarning false
	rows, err := db.Query("SELECT * FROM users WHERE id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

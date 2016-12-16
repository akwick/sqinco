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

	s := fmt.Sprintf("SELECT * FROM foo where name = '%s'", "Lovely Gopher")
	//	s := fmt.Sprintf("SELECT * FROM foo where name = '%s'", os.Args[1])

	// Fetching Data from a Database
	rows, err := db.Query("SELECT * FROM users WHERE id =" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

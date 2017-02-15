package sqlInjections

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func UnprepStmtConst() {
	// Accessing a DB
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer db.Close()

	// Fetching data from the database
	// @ExpectWarning false
	rows, err := db.Query("SELECT * FROM users WHERE id =" + "1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer rows.Close()
}

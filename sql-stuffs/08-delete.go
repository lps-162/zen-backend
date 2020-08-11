package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	database = "zentools"
	user     = "root"
	password = "sathyam"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	var connectionString = fmt.Sprintf("%s:%s@/%s", user, password, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	sqlStmt, err := db.Prepare("delete from products where name = ?")
	checkError(err)
	defer sqlStmt.Close()

	res, err := sqlStmt.Exec("IPhone")
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Println("Total Rows Deleted :", rowCount)

	fmt.Println()
}

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

	sqlStmt, err := db.Prepare("update products set quantity = ? where name = ?")
	checkError(err)
	defer sqlStmt.Close()

	res, err := sqlStmt.Exec(100, "IPhone")
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Println("Total Rows Updated :", rowCount)

	fmt.Println()
}

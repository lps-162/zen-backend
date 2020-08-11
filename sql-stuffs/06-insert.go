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

	sqlStmt, err := db.Prepare("insert into products(name, quantity) values (?, ?)")
	checkError(err)
	defer sqlStmt.Close()

	// res, err := sqlStmt.Exec("IPhone", 90)
	// checkError(err)
	// rowCount, err := res.RowsAffected()
	// fmt.Println("Inserted %d rows", rowCount)

	// res, err = sqlStmt.Exec("Android", 90)
	// checkError(err)
	// rowCount, err = res.RowsAffected()
	// fmt.Println("Inserted %d rows", rowCount)

	res, err := sqlStmt.Exec("Ipad", 90)
	checkError(err)
	newId, err := res.LastInsertId()
	fmt.Println("Inserted Id :", newId)

	fmt.Println()
}

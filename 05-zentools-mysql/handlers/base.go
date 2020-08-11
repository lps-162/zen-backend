package handlers

import "database/sql"

type Handler struct {
	DBConn *sql.DB
}

package routesHandlers

import "database/sql"

var Database *sql.DB

func SqlConnect() {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}
	Database = db
}

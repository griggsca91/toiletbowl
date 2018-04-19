package rename

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetTables() []string {
	db := GetDB()
	defer db.Close()

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}

	tables := make([]string, 0)
	for rows.Next() {
		var tableName string
		if err = rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, tableName)
	}

	return tables

}

func GetDB() *sql.DB {
	connStr := "postgresql://bruh@localhost:26257/logfather?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDB() {
	db := GetDB()
	defer db.Close()
	logTableSQL := `CREATE TABLE IF NOT EXISTS logfather.logs (
		id serial PRIMARY KEY, 
		message TEXT,
		file_name TEXT,
		line TEXT,
		function TEXT,
		created_on TIMESTAMP DEFAULT NOW()
	)`
	if _, err := db.Exec(logTableSQL); err != nil {
		log.Fatal(err)
	}

}

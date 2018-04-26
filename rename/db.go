package rename

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	databaseUser = "postgres"
	database     = "rename"
	password     = "password"
)

var (
	pool *sql.DB
)

func GetTables() []string {
	db := GetDB()

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
	if pool == nil {
		connStr := fmt.Sprintf("postgresql://%s@localhost:5432/%s"+
			"?sslmode=disable"+
			"&password=%s",
			databaseUser, database, password)
		pool, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	return pool
}

func InitDB() {
	db := GetDB()
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

package toiletbowl

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	databaseUser = "toiletbowl"
	database     = "toiletbowl"
	password     = "password"
)

var (
	pool *gorm.DB
)

func GetDB() *gorm.DB {
	if pool == nil {
		connStr := fmt.Sprintf(`sslmode=disable
			host=%s
			port=%s
			user=%s
			dbname=%s
			password=%s`,
			"localhost", "5432", databaseUser, database, password)
		var err error
		pool, err = gorm.Open("postgres", connStr)

		if err != nil {
			log.Fatal(err)
		}
	}

	return pool
}

func InitDB() {
	db := GetDB()
	db.CreateTable(&User{})
	db.CreateTable(&Poo{})

}

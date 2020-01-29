package main

import (
	"log"

	"../schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=sap dbname=postgres password='password' sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Connection Established")

	schema.Migrate(db)
}

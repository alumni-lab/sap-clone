package main

import (
	"log"
	"os"
	"strings"

	"../schema"
	"../seeds"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "sap")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "postgres")

	dbDialogue := createConnectionDialogue()

	db, err := gorm.Open("postgres", dbDialogue)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Connection Established")

	schema.Migrate(db)
	seeds.Seed(db)
}

func createConnectionDialogue() string {
	var dbDialogue strings.Builder

	dbDialogue.WriteString("host=")
	dbDialogue.WriteString(os.Getenv("DB_HOST"))
	dbDialogue.WriteString(" port=")
	dbDialogue.WriteString(os.Getenv("DB_PORT"))
	dbDialogue.WriteString(" user=")
	dbDialogue.WriteString(os.Getenv("DB_USER"))
	dbDialogue.WriteString(" dbname=")
	dbDialogue.WriteString(os.Getenv("DB_NAME"))
	dbDialogue.WriteString(" password=")
	dbDialogue.WriteString("'" + os.Getenv("DB_PASSWORD") + "'")
	dbDialogue.WriteString(" sslmode=disable")

	return dbDialogue.String()
}

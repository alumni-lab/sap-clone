package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"../../config"
	"../schema"
	"../seeds"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	config.RetrieveDBEnvVariables()

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
	dbDialogue.WriteString(os.Getenv("DBHost"))
	dbDialogue.WriteString(" port=")
	dbDialogue.WriteString(os.Getenv("DBPort"))
	dbDialogue.WriteString(" user=")
	dbDialogue.WriteString(os.Getenv("DBUser"))
	dbDialogue.WriteString(" dbname=")
	dbDialogue.WriteString(os.Getenv("DBName"))
	dbDialogue.WriteString(" password=")
	dbDialogue.WriteString("'" + os.Getenv("DBPassword") + "'")
	dbDialogue.WriteString(" sslmode=disable")

	fmt.Println(dbDialogue.String())

	return dbDialogue.String()
}

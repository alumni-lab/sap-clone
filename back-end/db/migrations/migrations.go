package main

import (
	"log"

	"github.com/alumni-lab/sap-clone/back-end/db/schema"
	"github.com/alumni-lab/sap-clone/back-end/db/seeds"
	"github.com/alumni-lab/sap-clone/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	config.RetrieveDBEnvVariables()

	dbDialogue := CreateConnectionDialogue()

	db, err := gorm.Open("postgres", dbDialogue)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Connection Established")

	schema.Migrate(db)
	seeds.Seed(db)
}

package main

import (
	"fmt"
	"log"

	"github.com/alumni-lab/sap-clone/config"
	"github.com/alumni-lab/sap-clone/db"
	"github.com/alumni-lab/sap-clone/db/schema"
	"github.com/alumni-lab/sap-clone/db/seeds"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	config.RetrieveDBEnvVariables()

	dbDialogue := db.CreateConnectionDialogue()

	db, err := gorm.Open("postgres", dbDialogue)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Connection Established")

	schema.Migrate(db)
	seeds.Seed(db)

	user := schema.Users{}

	_, error := user.FindUserByID(db, 1)
	if error != nil {
		panic(error)
	}

	fmt.Println(user)
}

package db

import (
	"fmt"
	"os"
	"strings"
)

// CreateConnectionDialogue creates string to connect to pgress db.
func CreateConnectionDialogue() string {
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

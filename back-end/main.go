package main

import "github.com/jinzhu/gorm"

func main() {
	db, err := gorm.Open("postgres", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()
}

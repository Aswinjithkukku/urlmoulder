package initializer

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("moulder.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully")
}

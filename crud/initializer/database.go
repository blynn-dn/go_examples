package initializer

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening db: %s", err)
	}

	db, err := DB.DB()
	if err != nil {
		log.Fatalf("Error accessing db: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}

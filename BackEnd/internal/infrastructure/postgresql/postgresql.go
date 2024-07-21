package postgresql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(postgres.Open("postgresql://neondb_owner:Z4zf7CFvglQK@ep-shrill-sound-a5g0r3g2.us-east-2.aws.neon.tech/dstock?sslmode=require"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get generic database object: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("failed to ping database: %w", err)
	}
	log.Println("Successfully connected to the database")
}

package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	host := "postgres"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "barang_db"

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB_HOST     : %s", host)
		log.Printf("DB_PORT     : %s", port)
		log.Printf("DB_USERNAME : %s", user)
		log.Printf("DB_NAME     : %s", dbname)

		fmt.Println("✅ Error to PostgreSQL with GORM")
		return nil, err
	}

	fmt.Println("✅ Connected to PostgreSQL with GORM")
	return db, nil
}

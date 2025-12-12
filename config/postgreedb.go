package config

import (
	"database/sql"
	"fmt"
	"os"
)

type PostgreedbDB struct {
	db *sql.DB
}

func NewDB(database *sql.DB) *PostgreedbDB {
	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")

	conststr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=service_api sslmode=disable", host, dbport, db_username, db_password)
	fmt.Println("Database connection string:", conststr)
	db, err := sql.Open("postgres", conststr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL database")

	return &PostgreedbDB{db: database}
}

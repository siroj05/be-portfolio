package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/siroj05/portfolio/database"
)

/*
* koneksi ke db disini
 */

var DB *sql.DB

func GetConnection() {
	// Kode untuk mendapatkan koneksi ke database
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using environmental variables.")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

	// koneksi ke db dengan retry
	var db *sql.DB
	for i := 0; i < 15; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Waiting for database connection... Attempt %d/15. Error: %v", i+1, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Println(err)
		log.Fatal("Could not connect to database after 15 attempts")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// Run migrations
	database.RunMigrations(db)

	DB = db
}


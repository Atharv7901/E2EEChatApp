package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load("/Users/atharvshivarkar/Desktop/projects/E2EEChatApp/backend/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configStruct := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 os.Getenv("DB_ADDRESS"),
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", configStruct.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func ConnectionTest(db *sql.DB){
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connection Successfull!")
}

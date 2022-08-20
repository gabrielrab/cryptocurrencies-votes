package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseConnection(){
	connectionString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s sslmode=disable password=%s", os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_DB"), os.Getenv("POSTGRESQL_PASS"))

	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	RunMigrations(db)
}

func CloseConnection() error {
	config, err := db.DB()
	if err != nil {
		return err
	}
	err = config.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
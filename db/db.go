package db

import (
	"fmt"
	"log"
	"os"
	"sample-api/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	user := "postgres"
	password := "postgrespw"
	host := "localhost"
	port := "55000"
	database := "todos"

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.Migrator().HasTable(&models.Task{}) {
		err := db.Migrator().CreateTable(&models.Task{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.Task{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

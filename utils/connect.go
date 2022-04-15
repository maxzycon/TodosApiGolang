package utils

import (
	"fmt"
	"log"
	"todosAPI/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=oskar101 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Makassar"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("conenction database success")
	DB = db
	DB.AutoMigrate(&database.Category{}, &database.Todos{})
	fmt.Println("Database Migrated")
}
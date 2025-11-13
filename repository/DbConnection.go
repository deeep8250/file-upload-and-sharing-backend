package repository

import (
	"file_uploading/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbCon() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("credential load failed")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection failed : %v", err.Error())
	}

	err = db.AutoMigrate(models.File{})
	if err != nil {
		log.Println("migration issue : ", err.Error())
		return
	}
}

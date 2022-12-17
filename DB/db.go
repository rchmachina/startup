package controller
import (
		"log"
		"github.com/joho/godotenv"
		"gorm.io/driver/postgres"
		"os"
		"gorm.io/gorm")


func Connect() *gorm.DB {

	err := godotenv.Load(".env")
	if err !=nil{
		log.Fatalf("error env")
	}


  db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
  if err != nil{
    panic(err)
  } 
  

	return db
}
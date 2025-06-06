package initializers

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
 
var DB *gorm.DB 

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(DB)
	if err != nil {
		log.Fatal("Error in connecting DB", err)
	}
}
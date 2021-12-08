package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type Gorm struct{}

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=localhost user=postgres dbname=cleancrud sslmode=disable password=postgres")

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}

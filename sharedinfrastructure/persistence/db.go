package persistence

import (
	"errors"
	"fmt"
	"os"
	"presentation/domain/entity"
	"presentation/domain/repository"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //postgres driver
)


type Repositories struct {
	Country repository.CountryRepository
	db   *gorm.DB
}
var Conn *gorm.DB

func ConnUsers() (Repositories, error) {
	fmt.Println("Attempting to connect to 'presentation' database...")
	err := godotenv.Load("config.env")
	if err != nil {
		err = errors.New("error accessing config file")
		fmt.Println(err)
	}

	username, password, dbName, dbHost, err := getDatabaseCredentials()
	if err != nil {
		fmt.Println(err)
	}

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		fmt.Println("postgres")
	}

	Conn = db
	fmt.Println("Core database connection successful")
	autoMigrateTables()
	return Repositories{
		Country: NewCountryInfra(db),
		db:   db,
	}, nil
}

func getDatabaseCredentials() (string, string, string, string, error) {
	_ = godotenv.Load("config.env")
	appMode := os.Getenv("app_mode")

	if appMode == "dev" {
		_ = godotenv.Load("dev-config.env")

		username := os.Getenv("db_user")
		password := os.Getenv("db_pass")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")

		return username, password, dbName, dbHost, nil
	}
	errorString := errors.New("error getting environment")
	return errorString.Error(), errorString.Error(), errorString.Error(), errorString.Error(), errorString
}

func autoMigrateTables() {
	Conn.AutoMigrate(&entity.CountryStruct{})
}

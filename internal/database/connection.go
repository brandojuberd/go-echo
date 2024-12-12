package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDatabaseConnection() *gorm.DB {
	err := godotenv.Load("development.env")
	if(err != nil){
		fmt.Println(err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(host, port, dbName, user, password)


	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     dbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}


	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})


	fmt.Println("Success connect/create to db: " + dbName)


	if err != nil {
		panic("failed to connect database")
	}
	return db
}

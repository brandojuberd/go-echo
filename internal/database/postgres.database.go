package database

import (
	"fmt"
	"go-echo/internal/config"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	dbInstance *postgresDatabase
)

func InitPostgresDatabase(dbConfig *config.Db) Database {
	user := dbConfig.User
	password := dbConfig.Password
	host := dbConfig.Host
	port := dbConfig.Port
	dbName := dbConfig.DBName

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
	dbInstance = &postgresDatabase{Db: db}

	return dbInstance
}

func (*postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Server struct {
		Port string
		JwtSecret string
	}

	Db struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Config struct {
		Server *Server
		Db     *Db
	}
)

func GetConfig() *Config {
	err := godotenv.Load("development.env")
	
	if err != nil {
		fmt.Println(err)
	}
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	serverPort := os.Getenv("PORT")
	serverJwtSecret := os.Getenv("JWT_SECRET")
	if serverJwtSecret == "" {
		serverJwtSecret = "secret"
	}

	config := &Config{
		Db: &Db{
			User:     user,
			Password: password,
			Host:     host,
			Port:     dbPort,
			DBName:   dbName,
		},
		Server: &Server{
			Port: serverPort,
			JwtSecret: serverJwtSecret,
		},
	}
	return config
}

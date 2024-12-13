package internal

import (
	"fmt"
	"go-echo/internal/database"
	"go-echo/internal/server"

	"github.com/joho/godotenv"
)

func InitInternal() {
	err := godotenv.Load("development.env")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Start go-echo Application")

	db := database.InitDatabaseConnection()

	echoServer := server.NewEchoServer(db)

	echoServer.Start()

}

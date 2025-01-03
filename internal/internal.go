package internal

import (
	"fmt"
	"go-echo/internal/config"
	"go-echo/internal/database"
	"go-echo/internal/server"
)

func InitInternal() {
	fmt.Println("Start go-echo Application")

	config := config.GetConfig()
	db := database.InitPostgresDatabase(config.Db)

	echoServer := server.NewEchoServer(db, *config.Server)

	echoServer.Start()
}

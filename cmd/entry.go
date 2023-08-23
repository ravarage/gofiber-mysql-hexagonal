package cmd

import (
	"fmt"
	"newglo/internals/database"
	"newglo/internals/handlers"
	"newglo/internals/server"
)

func Start() {
	db, err := database.New(&database.DatabaseConfig{
		Driver:   "nil",
		Host:     "127.0.0.1",
		Username: "nahry",
		Password: "402xsALYghA$",
		Port:     3306,
		Database: "nahry",
	})
	if err != nil {
		fmt.Print(err)
	}
	userHandlers := handlers.NewApp(db)
	//server
	httpServer := server.NewServer(
		userHandlers,
	)
	httpServer.Initialize()

}

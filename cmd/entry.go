package cmd

import (
	"fmt"
	"newglo/internals/database"
	"newglo/internals/handlers"
	"newglo/internals/repositories"
	"newglo/internals/server"
	"newglo/internals/services"
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
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			fmt.Print(err)
		}
	}(db)
	//
	repo := repositories.New(db.DB, db.Client, db.Context)

	userHandlers := services.NewUserService(repo)
	userRepository := handlers.NewApp(userHandlers)

	//userHandlers := handlers.NewApp(*db)
	//server
	httpServer := server.NewServer(
		userRepository,
	)
	httpServer.Initialize()

}

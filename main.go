package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"helpstack/config"
	"helpstack/config/database"
	"helpstack/server"
	"helpstack/utils"

	//"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("No .env file found")
	}
}


func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.("Error loading .env file")
	//}

	// Server initialization
	fiberApp := server.Create()
	conf:=config.New()
	Db, err:= dbPkg.NewSqlDb(&conf.SqlDbConfig)


	app:= server.MainApp{
		Engine: fiberApp,
		Configs: *conf,
	}


	// Migrations
	if err != nil {
		fmt.Println("failed to connect to database:", err.Error())
	} else {
		if Db == nil {
			fmt.Println("failed to connect to database: db variable is nil")
		} else {
			//fiberApp.DB = db
			utils.MigrateDb(Db)
		}
	}

	// Api routes
	server.SetupApiRoutes(fiberApp, dbPkg.Databases{Gorm: Db})

	if err := app.Listen(fiberApp); err != nil {
		log.Panic().Err(err).Msg("listen Error")
	}
}


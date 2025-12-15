package main

import (
	"auth-service/dbops"
	"auth-service/loggerconfig"
	"auth-service/router"

	"github.com/joho/godotenv"

	_ "auth-service/docs"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		// fmt.Print(e)
	}
	loggerconfig.InitLogrus()
	loggerconfig.Info("GIN auth-service started!")

	cfg, err := dbops.LoadConfig()
	if err != nil {
		loggerconfig.Panic("unable to load config")
	}
	err = dbops.InitPostgres(cfg)
	if err != nil {
		loggerconfig.Panic("Unable to connect db")
	}
	dbops.MigrateTables()
	r := router.InitRouters()
	port := "8080"
	r.Run(":" + port)
}

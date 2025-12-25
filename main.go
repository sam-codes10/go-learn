package main

import (
	"auth-service/dbops"
	"auth-service/kafka"
	"auth-service/loggerconfig"
	"auth-service/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "auth-service/docs"
)

func main() {
	env := os.Getenv("GO_ENV")
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	loggerconfig.InitLogrus()
	loggerconfig.Info("GIN auth-service started!")

	cfg, err := dbops.LoadConfig()
	if err != nil {
		loggerconfig.Panic("unable to load config")
	}

	err = dbops.InitPostgres(cfg, env)
	if err != nil {
		loggerconfig.Panic("Unable to connect db")
	}

	kafka.NewProducer(cfg, env, "otp")

	err = dbops.InitRedis(cfg, env)
	if err != nil {
		loggerconfig.Panic("Unable to connect redis")
	}

	dbops.MigrateTables()

	r := router.InitRouters()
	port := "8080"
	r.Run(":" + port)
}

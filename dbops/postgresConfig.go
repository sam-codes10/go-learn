package dbops

import (
	"auth-service/resourceConfig"
	"fmt"
	"log"

	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres(cfg resourceConfig.Config, env string) error {
	// env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	var pg resourceConfig.PostgresConfig
	// cfg, err := LoadConfig()
	// if err != nil {
	// 	return err
	// }

	switch env {
	case "local":
		pg = cfg.LocalConfig.Postgres
	default:
		pg = cfg.LocalConfig.Postgres
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg.Host, pg.Port, pg.User, pg.Password, pg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	DB = db

	return nil
}

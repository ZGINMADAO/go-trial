package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

type Database struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
}

var Db = Database{
	Driver: os.Getenv("DB_CONNECTION"),
	Host:   os.Getenv("DB_HOST"),
	Port:   os.Getenv("DB_PORT"),
	User:   os.Getenv("DB_USERNAME"),
	Pass:   os.Getenv("DB_PASSWORD"),
	Name:   os.Getenv("DB_DATABASE"),
}

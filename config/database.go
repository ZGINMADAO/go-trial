package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	Db = Database{
		Driver: os.Getenv("DB_CONNECTION"),
		Host:   os.Getenv("DB_HOST"),
		Port:   port,
		User:   os.Getenv("DB_USERNAME"),
		Pass:   os.Getenv("DB_PASSWORD"),
		Name:   os.Getenv("DB_DATABASE"),
	}

}

type Database struct {
	Driver string
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
}

var Db Database

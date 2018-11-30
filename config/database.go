package config

const DriverName = "mysql"

type Database struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

var Db = Database{
	Host: "106.12.16.28",
	Port: 3306,
	User: "root",
	Pass: "root",
	Name: "go-trial",
}

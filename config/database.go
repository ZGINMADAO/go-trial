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
	Host: "120.26.106.212",
	Port: 3306,
	User: "root",
	Pass: "W2lqsNVra86duEP9",
	Name: "go-trial",
}

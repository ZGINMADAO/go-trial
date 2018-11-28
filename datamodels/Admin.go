package datamodels

type Admin struct {
	Id       int
	Username string `xorm:"varchar(200)"`
	Password string `xrom:"varchar(32)"`
}

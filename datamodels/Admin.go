package datamodels

import "time"

type Admin struct {
	Id        int
	Avatar    string
	Nickname  string
	Username  string    `xorm:"'username' varchar(500)"`
	Password  string    `xrom:"char(32)"`
	CreatedAt time.Time `xorm:"datetime"`
	UpdatedAt time.Time `xorm:"datetime"`
}

func (_ *Admin) TableName() string {
	return "admin"
}

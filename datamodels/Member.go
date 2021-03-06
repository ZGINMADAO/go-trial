package datamodels

import "time"

type Member struct {
	Id        int
	Nickname  string
	Username  string
	Password  string
	Status    int8
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
}

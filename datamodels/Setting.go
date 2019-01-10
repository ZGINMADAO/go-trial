package datamodels

import "time"

type Setting struct {
	Id        int
	Key       string
	Val       string
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
}

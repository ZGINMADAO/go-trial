package datamodels

import "time"

type At struct {
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
}

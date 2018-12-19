package datamodels

import "time"

type Role struct {
	Id        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

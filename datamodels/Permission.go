package datamodels

import "time"

type Permission struct {
	Id        int
	Title     string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
}

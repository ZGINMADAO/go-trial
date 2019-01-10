package datamodels

import "time"

type Supplier struct {
	Id             int       `xorm:"pk autoincr notnull"`
	Title          string
	Sold           int
	Contact        string
	OpenAt         time.Time
	CloseAt        time.Time
	Address        string
	Describe       string
	BusinessStatus int8
	Score          string
	Status         int8
	CreatedAt      time.Time `xorm:"datetime created"`
	UpdatedAt      time.Time `xorm:"datetime updated"`
	//DeletedAt      time.Time `xorm:"datetime deleted"`
}

func (_ *Supplier) TableName() string {
	return "supplier"
}

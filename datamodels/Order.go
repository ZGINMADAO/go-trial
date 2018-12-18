package datamodels

import "time"

type Order struct {
	Id               int
	TradeNo          string    `xorm:"char(50) not null"`
	ProductId        int       `xorm:"int not null"`
	ProductTitle     string    `xorm:"char(50) not null"`
	ProductTypeId    int       `xorm:"int not null"`
	ProductTypeTitle string    `xorm:"char(50) not null"`
	OriginAmount     int64     `xorm:"bigint not null"`
	RealAmount       int64     `xorm:"bigint not null"`
	Status           int8      `xorm:"Int not null"`
	CreatedAt        time.Time `xorm:"datetime created"`
	UpdatedAt        time.Time `xorm:"datetime updated"`
	DeletedAt        time.Time `xorm:"datetime deleted"`
}

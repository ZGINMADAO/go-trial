package datamodels

import "time"

type Product struct {
	Id        int       `xorm:"pk autoincr notnull"` //凡是类型为 int64 且字段名为 Id 且没有定义任何 tag 的字段，都会自动被认为是主键
	TypeId    int       `xorm:"type_id INT" form:"typeId" validate:"gte=0"`
	Title     string    `xorm:"title CHAR(50) notnull" form:"title" validate:"required"`
	Price     int64     `xorm:"price BIGSERIAL" form:"price" validate:"gte=0"`
	Stock     int       `xorm:"stock INT" form:"stock" validate:"gte=0"`
	Detail    string    `xorm:"detail VARCHAR(200)" form:"detail"`
	Image     string    `xorm:"image VARCHAR(500)" form:"image"`
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
}

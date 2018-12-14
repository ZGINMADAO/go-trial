package datamodels

import "time"

type Product struct {
	Id        int       `xorm:"pk autoincr notnull"` //凡是类型为 int64 且字段名为 Id 且没有定义任何 tag 的字段，都会自动被认为是主键
	TypeId    int       `xorm:"type_id INT"`
	Title     string    `xorm:"title CHAR(50) notnull"`
	Price     int64     `xorm:"price BIGSERIAL"`
	Stock     int       `xorm:"stock INT"`
	Detail    string    `xorm:"detail VARCHAR(200)"`
	Image     string    `xorm:"image VARCHAR(500)"`
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
}

package datamodels

import "time"

type Tree struct {
	Id        int       `xorm:"pk autoincr notnull"` //凡是类型为 int64 且字段名为 Id 且没有定义任何 tag 的字段，都会自动被认为是主键
	Title     string    `xorm:"title char(50) notnull"`
	Path      string    `xorm:"varchar(500)"`
	Icon      string    `xorm:"char(50)"`
	ParentId  int       `xorm:"int"`
	CreatedAt time.Time `xorm:"datetime created"`
	UpdatedAt time.Time `xorm:"datetime updated"`
	DeletedAt time.Time `xorm:"datetime deleted"`
	Children  []*Tree   `xorm:"-"`
}

package datamodels

type Tree struct {
	Id int64 //凡是类型为 int64 且字段名为 Id 且没有定义任何 tag 的字段，都会自动被认为是主键
	Title string `xorm:""`


}
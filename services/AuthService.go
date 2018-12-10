package services

import (
	"go-trial/datamodels"
)


type authService struct {
}

func NewAuth() *authService {
	return &authService{}
}

type List struct {
	Title    string
	ParentId int
	Id       int
	Child    []*List
}

var one []List

//[]
func Recursive(Trees []datamodels.Tree, parentId int, node *List) interface{} {

	for _, val := range Trees {

		if parentId == val.ParentId {
			child := List{ParentId: val.ParentId, Id: val.Id, Child: []*List{}}
			node.Child = append(node.Child, &child)
			Recursive(Trees, val.ParentId, &child)

		}
	}
	return one
}

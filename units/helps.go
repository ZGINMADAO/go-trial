package units

import (
	"go-trial/datamodels"
)

type List struct {
	Id       int
	Title    string
	ParentId int
	Icon     string
	Path     string
	Children *[]List
}

func Recursive(Trees []datamodels.Tree, parentId int, nodeList *[]List) {
	for _, val := range Trees {

		if parentId == val.ParentId {

			temp := make([]List, 0)

			child := List{Id: val.Id, Title: val.Title, ParentId: val.ParentId, Icon: val.Icon, Path: val.Path, Children: &temp}

			*nodeList = append(*nodeList, child)

			Recursive(Trees, val.Id, &temp)
		}
	}
}

package backup

import (
	"fmt"
	"go-trial/datamodels"
)

type List struct {
	Id       int
	Title    string
	ParentId int
	Child    *[]List
}

func Recursive(Trees []datamodels.Tree, parentId int, nodeList *[]List) {
	for _, val := range Trees {

		if parentId == val.ParentId {

			temp := make([]List, 0)

			child := List{Id: val.Id, Title: val.Title, ParentId: val.ParentId, Child: &temp}

			*nodeList = append(*nodeList, child)

			fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", nodeList, nodeList, nodeList)

			Recursive(Trees, val.Id, &temp)
		}
	}
}

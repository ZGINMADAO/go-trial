package units

import (
	"go-trial/datamodels"
	"github.com/kataras/iris/core/errors"
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

func InArray(needle interface{}, haystack interface{}) (bool, error) {

	switch haystack.(type) {
	case []int:
		temp := haystack.([]int)
		for _, val := range temp {
			if needle == val {
				return true, nil

			}
			return false, nil
		}
	case []string:
		temp := haystack.([]string)
		for _, val := range temp {
			if needle == val {
				return true, nil

			}
			return false, nil
		}
	}
	return false, errors.New("请输入正确的格式")
}

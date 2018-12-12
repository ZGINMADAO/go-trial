package units

import (
	"go-trial/datamodels"
	"encoding/json"
	"fmt"
)

func Tree(list []*datamodels.Tree) string {
	data := BuildData(list)
	result := MakeTreeCore(0, data)

	fmt.Println("tree打印", result)

	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func BuildData(list []*datamodels.Tree) map[int]map[int]*datamodels.Tree {
	var data = make(map[int]map[int]*datamodels.Tree)
	for _, v := range list {
		id := v.Id
		fid := v.ParentId
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]*datamodels.Tree)
		}
		data[fid][id] = v
	}
	return data
}

func MakeTreeCore(index int, data map[int]map[int]*datamodels.Tree) []*datamodels.Tree {
	tmp := make([]*datamodels.Tree, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = MakeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}

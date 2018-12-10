package units

import (
	"go-trial/datamodels"
	"log"
)



//$trees = [
//[
//'id' => 1,
//'parent_id' => 0,
//'title' => '安徽'
//    ],
//[
//'id' => 2,
//'parent_id' => 1,
//'title' => '阜阳'
//    ],
//
//[
//'id' => 3,
//'parent_id' => 0,
//'title' => '浙江'
//    ],
//[
//'id' => 4,
//'parent_id' => 1,
//'title' => '合肥'
//    ],
//[
//'id' => 5,
//'parent_id' => 2,
//'title' => '阜南'
//    ]
//];
//
//function dump($data)
//{
//echo '<pre>';
//var_dump($data);
//echo '</pre>';
//
//}
//
//$deep = function ($parentId, &$results) use (&$deep, &$trees) {
//foreach ($trees as $key => $tree) {
//
//if ($tree['parent_id'] === $parentId) {
//
//$results[$key] = $tree;
//$results[$key]['child'] = [];
//$deep($tree['id'], $results[$key]['child']);
//}
//}
//};
//
//$results = [];
//$deep(0, $results);
//
//dump($results);

type List struct {
	Title    string
	ParentId int
	Id       int
	Child    []*List
}

func Recursive(Trees []datamodels.Tree, parentId int, nodeList []*List) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	i:=0
	for _, val := range Trees {

		if parentId == val.ParentId {

			child := List{ParentId: val.ParentId, Id: val.Id, Child: []*List{}}

			nodeList = append(nodeList, &child)

			Recursive(Trees, val.ParentId, nodeList[i].Child)
			i++
		}
	}
}

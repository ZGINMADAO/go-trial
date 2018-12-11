package units

import (
	"go-trial/datamodels"
	//"fmt"
	"fmt"
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
	Id       int
	Title    string
	ParentId int
	Child    []List
}

//func Recursive(Trees []datamodels.Tree, parentId int, nodeList *[]List) {
//
//	for _, val := range Trees {
//
//		if parentId == val.ParentId {
//
//			temp := make([]List, 0)
//
//			child := List{Id: val.Id, Title: val.Title, ParentId: val.ParentId, Child: &temp}
//
//			*nodeList = append(*nodeList, child)
//
//			fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", nodeList, nodeList, nodeList)
//
//			Recursive(Trees, val.Id, &temp)
//
//		}
//	}
//}



func Recursive(Trees []datamodels.Tree, parentId int, nodeList []List) {

	for _, val := range Trees {

		if parentId == val.ParentId {

			temp := make([]List, 0)

			child := List{Id: val.Id, Title: val.Title, ParentId: val.ParentId, Child: temp}

			tempNode := append(nodeList, child)

			fmt.Printf("addr of osa:%p  nodeList:%v", nodeList,  tempNode)
			fmt.Println()
			//
			//Recursive(Trees, val.Id, temp)

		}
	}
}


//func SolvePointer(nodeList *[]List) {
//	temp := *nodeList
//
//	for _, val := range temp {
//		if len(*val.Child) > 1 {
//			SolvePointer(val.Child)
//		}
//	}
//}

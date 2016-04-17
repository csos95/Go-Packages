package kdtree

import (
	"fmt"
)

func ExampleKDTree_Insert() {
	var tree = KDTree{nil, 2}
	var data = make([]interface{}, 2)
	data[0], data[1] = 5, 5
	tree.Insert("five", data)
	data[0], data[1] = 6, 6
	tree.Insert("six", data)
	data[0], data[1] = 4, 4
	tree.Insert("four", data)
	data[0], data[1] = 3, 3
	tree.Insert("three", data)
	data[0], data[1] = 7, 7
	tree.Insert("seven", data)
	data[0], data[1] = 10, 10
	tree.Insert("ten", data)
	fmt.Println(tree.Traverse("in"))
	//Output:
	// [three four five six seven ten]
}

func ExampleKDTree_Traverse() {
	var tree = KDTree{nil, 1}
	var data = make([]interface{}, 1)
	data[0] = 2
	tree.Insert(2, data)
	data[0] = 3
	tree.Insert(3, data)
	data[0] = 1
	tree.Insert(1, data)
	fmt.Println(tree.Traverse("in"))
	fmt.Println(tree.Traverse("pre"))
	fmt.Println(tree.Traverse("post"))
	//Output:
	// [1 2 3]
	// [2 1 3]
	// [1 3 2]
}

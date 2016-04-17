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
	// three four five six seven ten
}

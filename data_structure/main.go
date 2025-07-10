package main

import (
	"fmt"
	"sunfa/data_structure/binary_tree"
)

// 二叉树操作
func main() {
	n1 := binary_tree.NewTreeNode(nil, nil, 1)
	n2 := binary_tree.NewTreeNode(nil, nil, 2)
	n3 := binary_tree.NewTreeNode(nil, nil, 3)
	n4 := binary_tree.NewTreeNode(nil, nil, 4)
	n5 := binary_tree.NewTreeNode(nil, nil, 5)
	n6 := binary_tree.NewTreeNode(nil, nil, 6)
	n7 := binary_tree.NewTreeNode(nil, nil, 7)
	n8 := binary_tree.NewTreeNode(nil, nil, 8)
	n9 := binary_tree.NewTreeNode(nil, nil, 9)

	n1.AddLeft(n2)
	n1.AddRight(n3)
	n2.AddLeft(n4)
	n2.AddRight(n5)
	n3.AddLeft(n6)
	n3.AddRight(n7)
	n4.AddRight(n9)
	n5.AddLeft(n8)

	res := make([]int, 0)

	binary_tree.InOrder(n1, &res)
	fmt.Printf("values = %v", res)
}

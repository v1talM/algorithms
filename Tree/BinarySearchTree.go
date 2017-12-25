package main

import (
	"algorithm/Model"
	"fmt"
)
/**
 * @description: Binary Search Tree Exercise
 * @Note: BST sort & QuickSort make the same comparisons,
 *        but in a different order.
 */
func main() {
	root := Model.NewBinaryTreeNode(0)
	root.CreateBSTTree(9,5,3)
	tree := Model.NewBinaryTree(root)
	tree.InOrderWalk()
	fmt.Println(root)
}

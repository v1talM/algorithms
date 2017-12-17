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
	root.CreateBSTTree(8,4,2,1,3,5,10,9)
	tree := Model.NewBinaryTree(root)
	tree.InOrderWalk()
	node := Model.TreeSearchRecursive(root, 2)
	node2 := Model.TreeSearchIteration(root, 2)
	fmt.Println("\nresult:", node)
	fmt.Println("result:", node2)
	min := tree.Minimum()
	fmt.Println(min)
	max := tree.Maxmum()
	fmt.Println(max)
}
